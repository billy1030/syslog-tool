# Roadmap: Syslog Tool

**Project:** Syslog Tool
**Created:** 2026-04-23
**Phases:** 1

## Phase 1: Syslog Tool (Sender + Receiver)

**Goal:** Go UDP syslog tool with web UI — works as both receiver and sender in one binary. Configurable port. Real-time log display with split panels.

**Requirements covered:** SYSLOG-01, SYSLOG-02, SYSLOG-03, SYSLOG-04, UI-01, UI-02, SENDER-01, SENDER-02, SENDER-03

**Success criteria:**
1. UDP listener binds to configurable port (default 514)
2. Web browser displays received logs in upper panel
3. Connection info (IP, port, timestamp) shown in lower panel
4. Only last 20 logs retained in memory
5. Option to save logs to file available
6. Sender can send UDP syslog to configurable destination IP:port
7. Single binary handles both roles

---

## Traceability

| Requirement | Phase | Status |
|-------------|-------|--------|
| SYSLOG-01 | Phase 1 | Pending |
| SYSLOG-02 | Phase 1 | Pending |
| SYSLOG-03 | Phase 1 | Pending |
| SYSLOG-04 | Phase 1 | Pending |
| UI-01 | Phase 1 | Pending |
| UI-02 | Phase 1 | Pending |
| SENDER-01 | Phase 1 | Pending |
| SENDER-02 | Phase 1 | Pending |
| SENDER-03 | Phase 1 | Pending |

**Coverage:** 9 requirements | 1 phase
