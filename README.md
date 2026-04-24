# Simple Syslog Tool

A lightweight UDP syslog receiver with a real-time web interface.

## Download

Download the pre-built binary for your platform:

- **Windows:** `syslog-tool.exe`
- **Linux AMD64:** `syslog-tool-linux-amd64`
- **Linux ARM64:** `syslog-tool-linux-arm64`

## Usage

```bash
./syslog-tool.exe -port 514 -web-port 18080
```

### Options

- `-port`       UDP port to listen for syslog (default: 514)
- `-web-port`   HTTP port for web UI (default: 18080)
- `-H`          Show help

### Examples

```bash
./syslog-tool.exe                  # defaults: UDP 514, web :18080
./syslog-tool.exe -port 514         # custom UDP port
./syslog-tool.exe -web-port 18080   # custom web port
./syslog-tool.exe -H                # show help
```

## Linux Setup

Ports below 1024 require root privileges:

```bash
chmod +x syslog-tool-linux-amd64
sudo ./syslog-tool-linux-amd64 -port 514 -web-port 18080
```

## Web UI

Open `http://localhost:18080` in your browser.

- **Received Logs** — real-time log stream via WebSocket
- **Connection Info** — listener address, total received, last update
- **Save** — download logs as `syslog.txt`
- **Exit** — stop the program
- **Clear** — clear the log display

## Features

- UDP syslog receiver (RFC 3164)
- Real-time WebSocket push to browser
- Circular buffer (last 65535 entries)
- Send syslog messages from UI
- Auto-save logs to file
- Grey `<34>` priority delimiter display
- Built by Billy Lam with AI assistance (2026)
