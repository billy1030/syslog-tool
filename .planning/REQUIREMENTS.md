# Requirements: Syslog Tool

**Defined:** 2026-04-23
**Core Value:** Real-time syslog monitoring with a clean, minimal web interface — zero configuration to start receiving logs.

## v1 Requirements

### Syslog Receiver

- [ ] **SYSLOG-01**: Listen for syslog messages on UDP port 514
- [ ] **SYSLOG-02**: Parse incoming syslog packets and extract message content
- [ ] **SYSLOG-03**: Retain only the last 20 received logs in memory
- [ ] **SYSLOG-04**: Provide option to save received logs to a file

### Web UI

- [ ] **UI-01**: Display received logs in upper panel (last 20, newest at top)
- [ ] **UI-02**: Display connection info in lower panel (source IP, port, received timestamp)

### Syslog Sender

- [ ] **SENDER-01**: Windows CLI tool to send syslog messages
- [ ] **SENDER-02**: Support configurable destination IP and port
- [ ] **SENDER-03**: Allow custom message text and basic formatting

## Out of Scope

| Feature | Reason |
|---------|--------|
| TCP syslog | UDP is the most common use case |
| Database persistence | Live display only |
| Authentication | Internal dev tool, no auth needed |
| Multi-threaded/clustered receiver | Single instance sufficient for debugging |
| macOS sender | Windows first, Linux receiver next |

---

*Requirements defined: 2026-04-23*
*Last updated: 2026-04-23 after initial definition*
