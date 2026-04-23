# Requirements: Syslog Tool

**Defined:** 2026-04-23
**Core Value:** Real-time syslog monitoring with a clean, minimal web interface — zero configuration to start receiving logs.

## v1 Requirements

### Syslog Receiver

- [ ] **SYSLOG-01**: Listen for syslog messages on configurable UDP port (default 514)
- [ ] **SYSLOG-02**: Parse incoming syslog packets and extract message content
- [ ] **SYSLOG-03**: Retain only the last 20 received logs in memory
- [ ] **SYSLOG-04**: Provide option to save received logs to a file

### Web UI

- [ ] **UI-01**: Display received logs in upper panel (last 20, newest at top)
- [ ] **UI-02**: Display connection info in lower panel (source IP, port, received timestamp)

### Syslog Sender

- [ ] **SENDER-01**: Send syslog messages via UDP to configurable destination IP:port
- [ ] **SENDER-02**: Allow custom message text input
- [ ] **SENDER-03**: Single binary handles both sender and receiver roles

## Out of Scope

| Feature | Reason |
|---------|--------|
| TCP syslog | UDP is the most common use case |
| Database persistence | Live display only |
| Authentication | Internal dev tool, no auth needed |
| Multi-threaded/clustered | Single instance sufficient for debugging |

---

*Requirements defined: 2026-04-23*
*Last updated: 2026-04-23 after scope update (combined sender+receiver into one phase)*
