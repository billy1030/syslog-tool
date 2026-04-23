# Phase 1: Syslog Receiver - Discussion Log

> **Audit trail only.** Do not use as input to planning, research, or execution agents.
> Decisions are captured in CONTEXT.md — this log preserves the alternatives considered.

**Date:** 2026-04-23
**Phase:** 1-syslog-receiver
**Areas discussed:** UI Framework, Go Libraries

---

## UI Framework

| Option | Description | Selected |
|--------|-------------|----------|
| Embed HTML/JS directly | Single binary, Go embeds assets via embed package | ✓ |
| Go templates | Server-side rendered HTML | |
| Separate frontend | React/Vue with separate build pipeline | |

**User's choice:** Embed HTML/JS directly
**Notes:** Recommended approach — single binary distribution, no frontend build pipeline needed

---

## Go Libraries

| Option | Description | Selected |
|--------|-------------|----------|
| gorilla/websocket | WebSocket support | ✓ |
| Standard library UDP | net.ListenPacket for UDP, no extra dependency | ✓ |
| Go embed | Bundle static assets into binary | ✓ |

**User's choice:** Use recommended stack
**Notes:** Simple and minimal approach

---

## Claude's Discretion

- Exact CSS styling (colors, fonts, layout spacing)
- WebSocket reconnection strategy on browser side
- Log format parsing (handle syslog RFC 3164 vs RFC 5424)

