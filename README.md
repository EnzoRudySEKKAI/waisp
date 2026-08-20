# waïsp

## The collaboration infrastructure for personal AI assistants

waïsp is a self-hosted A2A (Agent-to-Agent) platform that connects personal AI assistants inside enterprise environments.

Instead of providing one shared company chatbot, waïsp gives every person their own AI assistant that can:

- understand its user
- interact with the local machine
- access authorized company knowledge
- communicate with other assistants
- coordinate work across teams

The result is a network of AI representatives working together with humans.

---

# Vision

Modern companies will not only have humans collaborating through applications. They will have humans collaborating through intelligent assistants.

waïsp provides the infrastructure layer for this future:

```
User
 |
Personal Assistant
 |
Local Harness
 |
Company Brain
 |
Other Assistants
```

---

# Main Concepts

## Personal Assistant

Every user owns exactly one assistant.

The assistant contains personal context:

- preferences
- language
- working style
- personal memory
- private knowledge

It acts as the user's representative inside the organization.

---

## Company Brain

Each company runs its own self-hosted Brain.

The Brain is responsible for:

- assistant discovery
- identity management
- permissions
- conversations
- logs
- collective memory
- enterprise RAG

The Brain is the shared intelligence layer of the organization.

---

## Domains

Companies can create domains and subdomains:

```
Company
├── Engineering
│   ├── Backend
│   └── Frontend
├── Product
└── Marketing
```

Domains define:

- teams
- responsibilities
- permissions
- available context

The knowledge remains centralized in the Company Brain and is filtered according to access rules.

---

# Harness

waïsp uses a terminal-first harness approach inspired by tools like Claude Code and OpenCode.

The harness provides:

- terminal interface
- AI assistant interaction
- local machine access
- command execution
- file operations

The first implementation integrates a DeepSeek-based harness instead of rebuilding an execution engine from scratch.

---

# Architecture Direction

Initial stack:

- Go core services
- TUI applications
- A2A protocol communication
- Self-hosted Company Brain
- Modular plugin-based architecture

The architecture is designed to support future extensions:

- additional harness providers
- new AI models
- enterprise tools
- workflow automation

---

# Roadmap

## Phase 1

Foundation:

- Company Brain
- personal assistants
- TUI harness
- A2A communication
- basic permissions

## Phase 2

Knowledge and collaboration:

- enterprise RAG
- memory system
- advanced permissions
- conversation history

## Phase 3

Enterprise automation:

- task management
- calendars
- external integrations
- autonomous workflows

---

# Documentation

- [Vision](docs/vision.md)

---

# Status

waïsp is currently in early architecture and foundation development.
