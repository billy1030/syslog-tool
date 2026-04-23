package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
)

//go:embed assets
var assetsFS embed.FS

// LogEntry represents a single syslog message
type LogEntry struct {
	Timestamp  time.Time `json:"timestamp"`
	SourceIP   string    `json:"source_ip"`
	SourcePort int       `json:"source_port"`
	Message    string    `json:"message"`
}

// CircularBuffer holds the last N log entries
type CircularBuffer struct {
	entries [20]LogEntry
	count   int
	head    int
	mu      sync.Mutex
}

// Add adds a new entry to the buffer
func (cb *CircularBuffer) Add(entry LogEntry) {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.entries[cb.head] = entry
	cb.head = (cb.head + 1) % 20
	if cb.count < 20 {
		cb.count++
	}
}

// GetAll returns all entries newest first
func (cb *CircularBuffer) GetAll() []LogEntry {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	result := make([]LogEntry, 0, cb.count)
	for i := 0; i < cb.count; i++ {
		idx := (cb.head - cb.count + i + 20) % 20
		result = append(result, cb.entries[idx])
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

// Hub manages WebSocket clients
type Hub struct {
	clients    map[*websocket.Conn]bool
	broadcast  chan LogEntry
	register   chan *websocket.Conn
	unregister chan *websocket.Conn
	mu         sync.RWMutex
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client] = true
			h.mu.Unlock()
			for _, entry := range logBuffer.GetAll() {
				data, _ := json.Marshal(entry)
				client.WriteMessage(websocket.TextMessage, data)
			}
		case client := <-h.unregister:
			h.mu.Lock()
			delete(h.clients, client)
			h.mu.Unlock()
			client.Close()
		case entry := <-h.broadcast:
			data, _ := json.Marshal(entry)
			h.mu.RLock()
			for client := range h.clients {
				client.WriteMessage(websocket.TextMessage, data)
			}
			h.mu.RUnlock()
		}
	}
}

var (
	hub         = Hub{
		clients:    make(map[*websocket.Conn]bool),
		broadcast:  make(chan LogEntry, 100),
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
	}
	logBuffer  = CircularBuffer{}
	upgrader   = websocket.Upgrader{}
	listenPort int
)

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}
	hub.register <- conn
	defer func() {
		hub.unregister <- conn
	}()
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}
}

func handleSaveUI(w http.ResponseWriter, r *http.Request) {
	entries := logBuffer.GetAll()
	var sb strings.Builder
	for _, e := range entries {
		sb.WriteString(fmt.Sprintf("[%s] %s:%d %s\n",
			e.Timestamp.Format("2006-01-02 15:04:05"),
			e.SourceIP, e.SourcePort, e.Message))
	}
	w.Header().Set("Content-Disposition", "attachment; filename=syslog.txt")
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(sb.String()))
}

type SendRequest struct {
	DestIP   string `json:"dest_ip"`
	DestPort int    `json:"dest_port"`
	Message  string `json:"message"`
}

func handleSend(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST only", http.StatusMethodNotAllowed)
		return
	}
	var req SendRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.ParseIP(req.DestIP),
		Port: req.DestPort,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	msg := fmt.Sprintf("<34>%s", req.Message)
	conn.Write([]byte(msg))
	w.WriteHeader(http.StatusOK)
}

func handleAPIInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	entries := logBuffer.GetAll()
	info := map[string]interface{}{
		"listening_port": listenPort,
		"total_received": len(entries),
	}
	if len(entries) > 0 {
		info["last_update"] = entries[0].Timestamp.Format(time.RFC3339)
	}
	json.NewEncoder(w).Encode(info)
}

func main() {
	listenPort = 514
	if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-port=") {
		fmt.Sscanf(os.Args[1], "-port=%d", &listenPort)
	} else if len(os.Args) > 2 && os.Args[1] == "-port" {
		fmt.Sscanf(os.Args[2], "%d", &listenPort)
	}

	go hub.Run()

	addr := fmt.Sprintf(":%d", listenPort)
	conn, err := net.ListenPacket("udp", addr)
	if err != nil {
		log.Fatalf("Failed to bind UDP port %d: %v", listenPort, err)
	}
	defer conn.Close()

	log.Printf("Syslog receiver listening on UDP %d", listenPort)
	log.Printf("Web UI available at http://localhost:8080")

	go func() {
		buffer := make([]byte, 65535)
		for {
			n, src, err := conn.ReadFrom(buffer)
			if err != nil {
				continue
			}
			entry := LogEntry{
				Timestamp: time.Now(),
				SourceIP:  src.String(),
				Message:   string(buffer[:n]),
			}
			if udpAddr, ok := src.(*net.UDPAddr); ok {
				entry.SourcePort = udpAddr.Port
			}
			logBuffer.Add(entry)
			hub.broadcast <- entry
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/ws", handleWebSocket)
	mux.HandleFunc("/save", handleSaveUI)
	mux.HandleFunc("/api/send", handleSend)
	mux.HandleFunc("/api/info", handleAPIInfo)
	mux.Handle("/", http.FileServer(http.FS(assetsFS)))

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Println("Shutting down...")
		os.Exit(0)
	}()

	http.ListenAndServe(":8080", mux)
}
