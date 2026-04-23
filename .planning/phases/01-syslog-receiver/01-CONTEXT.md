# Phase 1: Syslog Tool - Context

**Gathered:** 2026-04-23
**Status:** Ready for planning

<domain>
## Phase Boundary

Go UDP syslog tool that works as both receiver and sender in one binary. Configurable listening port (default 514). Split-panel web UI — upper panel shows last 20 received logs, lower panel shows connection info (source IP, port, timestamp). Logs can be saved to file optionally. Single binary distribution.

</domain>

<decisions>
## Implementation Decisions

### Dual Role
- **D-01:** Single binary handles both send and receive roles
- **D-02:** Mode determined by command-line flags or UI toggle (not both simultaneously)

### Port Configuration
- **D-03:** Default port is 514 (standard syslog)
- **D-04:** Port configurable via command-line flag (e.g., `-port 9514`)
- **D-05:** Port shown in UI connection info panel

### UI Framework
- **D-06:** Embed HTML/JS directly — Go embeds `assets/` folder via `embed` package, single binary, no frontend build pipeline
- **D-07:** WebSocket for real-time log push — browser receives new logs without polling

### Go Libraries
- **D-08:** Use `github.com/gorilla/websocket` for WebSocket support
- **D-09:** Standard library for UDP (`net.ListenPacket`) — no third-party dependency needed
- **D-10:** `embed` package for bundling static assets into binary

### Log Display
- **D-11:** Circular buffer of 20 logs — newest at top, oldest discarded
- **D-12:** Each log entry shows: timestamp, source IP, source port, message

### Connection Info Panel
- **D-13:** Shows: listener address (IP:port), total logs received, last update time

### Sender Feature
- **D-14:** Sender sends UDP syslog messages to configurable destination IP:port
- **D-15:** Message text input in UI or via CLI flag

### File Save
- **D-16:** Optional — triggered via button in UI or command-line flag
- **D-17:** Save format: one line per log, text format with timestamp and source

### Claude's Discretion
- Exact CSS styling (colors, fonts, layout spacing)
- WebSocket reconnection strategy on browser side
- Log format parsing (handle syslog RFC 3164 vs RFC 5424)
- How sender destination is configured (UI input or CLI flag)

</decisions>

<canonical_refs>
## Canonical References

**Downstream agents MUST read these before planning or implementing.**

No external specs — requirements fully captured in decisions above.

### Syslog Reference
- RFC 3164 (BSD syslog) — format expected from most devices
- RFC 5424 (newer syslog format) — should also be parseable

</canonical_refs>

<code_context>
## Existing Code Insights

### Reusable Assets
- None yet — fresh project

### Established Patterns
- None yet — fresh project

### Integration Points
- UDP listener → WebSocket hub → browser
- UDP sender ← UI input
- Single `main.go` entry point with mode routing

</code_context>

<specifics>
## Specific Ideas

- "Simple and minimal" — don't overcomplicate the UI
- Single executable is important for easy distribution
- Sender is for testing the receiver

</specifics>

<deferred>
## Deferred Ideas

None — discussion stayed within phase scope

</deferred>

---

*Phase: 01-syslog-tool*
*Context gathered: 2026-04-23*
