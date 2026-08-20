# waïsp — Vision

## Overview

waïsp is a self-hosted A2A (Agent-to-Agent) collaboration infrastructure designed to connect personal AI assistants inside enterprise environments.

The goal is to create a network of personal assistants where every person owns an assistant capable of understanding their context, interacting with their local environment, and collaborating with other assistants.

Each assistant communicates through an A2A protocol and is coordinated by a company-level Collaboration Brain.

---

# Core Vision

The future of enterprise collaboration will involve humans working with personal AI representatives.

waïsp enables assistants to:

- understand their users
- execute local actions
- communicate with other assistants
- access company knowledge
- coordinate work
- preserve organizational context

---

# Personal Assistant

## One Person = One Assistant

Each user owns one personal assistant representing them inside the organization.

The assistant maintains:

- preferences
- communication style
- languages
- working habits
- personal knowledge
- private context

---

# Memory Model

waïsp uses two memory layers.

## Personal Memory

Private memory belonging to the user:

- preferences
- habits
- private information
- personal history

This memory follows the user everywhere.

## Domain Memory

Knowledge belonging to an organization context:

- company
- teams
- projects
- documentation
- decisions
- responsibilities

When changing domains, the assistant keeps personal memory but switches domain context.

---

# Company Brain

Each company runs its own self-hosted Collaboration Brain.

The Brain manages:

- identities
- assistants
- domains
- permissions
- conversations
- logs
- collective memory
- RAG knowledge bases

The Brain is the shared intelligence layer of the organization.

---

# Domains

A company can create multiple domains and subdomains.

Example:

```
Company
├── Engineering
│   ├── Backend
│   └── Frontend
├── Product
└── Marketing
```

Each domain has its own:

- members
- permissions
- memory
- documents
- RAG context
- conversations

---

# Harness

The assistant runs through a local harness inspired by tools such as Claude Code.

The harness provides:

- terminal interaction
- AI assistant interface
- local machine access
- file operations
- command execution
- context awareness

The goal is: one harness to interact with the assistant and accomplish work.

The first implementation will use a DeepSeek-based harness approach, with the possibility of supporting additional harness providers later.

---

# A2A Collaboration

Assistants communicate using an existing A2A protocol.

Flow:

```
Assistant A
    |
 A2A Protocol
    |
Company Brain
    |
Assistant B
```

The Brain handles:

- discovery
- permissions
- history
- context

If the Brain is unavailable, assistants can temporarily use cached permissions and synchronize later.

---

# Deployment

## Phase 1: LAN

The Brain runs inside the company's local network.

Benefits:

- privacy
- data control
- enterprise adoption

## Phase 2: Online

The Brain becomes accessible through the internet for distributed organizations.

---

# Long-Term Goal

waïsp aims to become the collaboration infrastructure where humans and AI assistants work together.

Every person has a trusted AI representative connected to a shared company intelligence.
