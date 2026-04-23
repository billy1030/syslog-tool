# Phase 1: Syslog Receiver - Context

**Gathered:** 2026-04-23
**Status:** Ready for planning

<domain>
## Phase Boundary

Go UDP listener on port 514 with a split-panel web UI. Upper panel shows last 20 received log messages (real-time). Lower panel shows connection info (source IP, port, timestamp). Logs can be saved to file optionally. Single binary distribution.

</domain>

<decisions>
## Implementation Decisions

### UI Framework
- **D-01:** Embed HTML/JS directly — Go embeds `assets/` folder via `embed` package, single binary, no frontend build pipeline
- **D-02:** WebSocket for real-time log push — browser receives new logs without polling

### Go Libraries
- **D-03:** Use `github.com/gorilla/websocket` for WebSocket support
- **D-04:** Standard library for UDP (`net.ListenPacket`) — no third-party dependency needed
- **D-05:** `embed` package for bundling static assets into binary

### Log Display
- **D-06:** Circular buffer of 20 logs — newest at top, oldest discarded
- **D-07:** Each log entry shows: timestamp, source IP, source port, message

### Connection Info Panel
- **D-08:** Shows: listener address (IP:port), total logs received, last update time

### File Save
- **D-09:** Optional — triggered via button in UI or command-line flag
- **D-10:** Save format: one line per log, text format with timestamp and source

### Claude's Discretion
- Exact CSS styling (colors, fonts, layout spacing)
- WebSocket reconnection strategy on browser side
- Log format parsing (handle syslog RFC 3164 vs RFC 5424)

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
- Single `main.go` entry point

</code_context>

<specifics>
## Specific Ideas

- "Simple and minimal" — don't overcomplicate the UI
- Single executable is important for easy distribution

</specifics>

<deferred>
## Deferred Ideas

None — discussion stayed within phase scope

</deferred>

---

*Phase: 01-syslog-receiver*
*Context gathered: 2026-04-23*
