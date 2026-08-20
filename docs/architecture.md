# waïsp Architecture

## Overview

waïsp is built around three major components:

```
User
 |
Harness (local TUI)
 |
Personal Assistant Runtime
 |
A2A Protocol
 |
Company Brain (self-hosted)
```

## Components

### Harness
Local execution environment inspired by Claude Code.

Responsibilities:
- terminal interface
- local machine interaction
- tool execution
- user interaction
- assistant access

Initial implementation integrates DeepSeek Harness.

### Personal Assistant Runtime
Represents one user.

Responsibilities:
- manage personal memory
- communicate through A2A
- decide actions
- interact with harness

### Company Brain
Self-hosted enterprise server.

Responsibilities:
- identity registry
- permissions
- domains
- conversations
- logs
- company memory
- RAG retrieval

## Design Principles

- modular components
- explicit interfaces
- replaceable providers
- minimal core
- plugin-oriented architecture
