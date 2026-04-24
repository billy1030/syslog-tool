# Simple Syslog Tool

A lightweight UDP syslog receiver with a real-time web interface.

## Usage

```bash
syslog-tool.exe -port 514 -web-port 18080
```

### Options

- `-port`       UDP port to listen for syslog (default: 514)
- `-web-port`   HTTP port for web UI (default: 18080)
- `-H`          Show help

### Examples

```bash
./syslog-tool.exe                  # defaults: UDP 514, web :18080
./syslog-tool.exe -port 514         # custom UDP port
./syslog-tool.exe -web-port 8080    # custom web port
./syslog-tool.exe -H                # show help
```

## Web UI

Open `http://localhost:18080` in your browser.

- **Received Logs** — real-time log stream via WebSocket
- **Connection Info** — listener address, total received, last update
- **Save** — download logs as `syslog.txt`
- **Exit** — stop the program
- **Clear** — clear the log display

## Building

```bash
# Windows
go build -o syslog-tool.exe

# Linux
GOOS=linux GOARCH=amd64 go build -o syslog-tool-linux
```

## Features

- UDP syslog receiver (RFC 3164)
- Real-time WebSocket push to browser
- Circular buffer (last 20 entries)
- Send syslog messages from UI
- Auto-save logs to file
- Grey `<34>` priority delimiter display
- Built by Billy Lam with AI assistance (2026)
