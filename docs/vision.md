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

waïsp separates personal memory from company knowledge.

## Personal Memory

Private memory belonging to the user:

- preferences
- habits
- private information
- personal history
- individual working context

This memory follows the user everywhere.

## Company Knowledge Memory

The shared knowledge belongs to the Company Brain.

The Company Brain contains a centralized RAG system containing enterprise knowledge while enforcing access restrictions based on:

- user permissions
- domain membership
- team access
- project access
- security rules

The RAG is not duplicated per domain. Instead, the Brain manages one global knowledge layer with controlled retrieval depending on the user's authorized context.

When changing domains, the assistant keeps personal memory but receives different accessible knowledge depending on the active domain and permissions.

---

# Company Brain

Each company runs its own self-hosted Collaboration Brain.

The Brain is the central intelligence layer of the organization.

It manages:

- identities
- assistants
- domains
- permissions
- conversations
- logs
- collective memory
- centralized RAG knowledge base

The Brain provides a shared context while maintaining strict information boundaries between users and teams.

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

Domains define organizational structure and access rules.

Each domain has:

- members
- permissions
- accessible knowledge
- documents
- conversations
- context

The knowledge itself remains stored in the Company Brain RAG, while retrieval is filtered according to domain and user permissions.

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

- assistant discovery
- permission validation
- conversation history
- context routing
- communication logs

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
