# Syslog Tool

## What This Is

A cross-platform Go syslog tool that works as both receiver and sender in one binary. Configurable listening port (default 514). Web browser UI split into two panels: upper panel shows last 20 received log messages, lower panel shows connection info (source IP, port, timestamp). Logs can be saved to file optionally.

## Core Value

Real-time syslog monitoring with a clean, minimal web interface — zero configuration to start receiving logs.

## Requirements

### Active

- [ ] Syslog receiver listens on configurable UDP port
- [ ] Web UI shows last 20 log messages (upper panel)
- [ ] Web UI shows connection info — source IP, port, received time (lower panel)
- [ ] Option to save logs to a file
- [ ] Syslog sender to send test messages to a destination IP:port

## Context

- Platform: Windows and Linux
- Use case: debugging/dev tool for syslog-emitting services
- No persistence required — live display only with optional file save

## Constraints

- **Protocol**: UDP only
- **Default port**: 514 (standard syslog), but configurable
- **Cross-platform**: Must work on Windows and Linux
- **No database**: Logs are in-memory only, file save is optional

## Key Decisions

| Decision | Rationale | Outcome |
|----------|-----------|---------|
| UDP only | Standard syslog, most common | — Pending |
| Go | Cross-platform, good for network tools | — Pending |
| Web UI | Clean, accessible, cross-platform | — Pending |
| Combined sender+receiver | Single binary, both roles | — Pending |
| No database | Keep it simple — live display only | — Pending |

---
*Last updated: 2026-04-23 after scope update (combined sender+receiver)*
