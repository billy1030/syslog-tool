# Syslog Tool

## What This Is

A cross-platform Go syslog receiver with a web browser UI for viewing received logs in real-time. The UI is split into two panels: upper panel shows the last 20 received log messages, lower panel shows connection information (source IP, port, etc.). Includes an optional log file save feature and a companion sender tool for testing.

## Core Value

Real-time syslog monitoring with a clean, minimal web interface — zero configuration to start receiving logs.

## Requirements

### Active

- [ ] Syslog receiver listens on UDP 514
- [ ] Web UI shows last 20 log messages (upper panel)
- [ ] Web UI shows connection info — source IP, port, received time (lower panel)
- [ ] Option to save logs to a file
- [ ] Companion syslog sender tool (Windows) for testing

## Context

- Platform: Windows first, then Linux
- Use case: debugging/dev tool for syslog-emitting services
- No persistence required — live display only with optional file save

## Constraints

- **Protocol**: UDP only (port 514)
- **Cross-platform**: Must work on Windows and Linux
- **No database**: Logs are in-memory only, file save is optional

## Key Decisions

| Decision | Rationale | Outcome |
|----------|-----------|---------|
| UDP 514 | Standard syslog port, most common | — Pending |
| Go for receiver | Cross-platform, good for network tools | — Pending |
| Web UI | Clean, accessible, cross-platform | — Pending |
| No database | Keep it simple — live display only | — Pending |

---
*Last updated: 2026-04-23 after initialization*
