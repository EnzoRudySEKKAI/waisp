# waïsp — Vision

## Overview

waïsp is a self-hosted A2A (Agent-to-Agent) collaboration infrastructure that enables every person inside an organization to have a personal AI assistant capable of acting locally and collaborating with other assistants.

The goal is not to build another company chatbot. The goal is to create a network of personal AI representatives connected through an enterprise Collaboration Brain.

Each person owns one assistant. The assistant understands its user's context, interacts with the local machine through a harness, and communicates with other assistants when collaboration is required.

---

# The Problem

Current enterprise AI tools are mainly centralized assistants:

- one chatbot for everyone
- limited personal context
- weak collaboration between users
- isolated knowledge sources
- no autonomous communication between AI agents

waïsp introduces a different model:

> Every employee has a trusted AI representative connected to the collective intelligence of the organization.

---

# Core Principles

## One Person = One Assistant

Every user has exactly one personal assistant identity.

The assistant represents the user and maintains:

- personal preferences
- communication style
- preferred languages
- working habits
- private knowledge
- personal history

The assistant follows the user across different projects and domains.

---

# Memory Architecture

waïsp separates personal intelligence from company intelligence.

## Personal Memory

Private user memory:

- preferences
- habits
- private notes
- personal workflows
- individual context

This memory belongs only to the user.

## Company Knowledge Brain

Each company owns one Collaboration Brain containing a centralized knowledge system.

The Brain includes a global RAG layer containing enterprise knowledge.

The RAG is not duplicated by domain. Instead, access is controlled dynamically through:

- user permissions
- domain membership
- team permissions
- project access
- security policies

The assistant retrieves only information that the user is authorized to access.

---

# Company Brain

Each organization runs its own self-hosted Company Brain.

The Brain is the coordination and knowledge layer of the company.

Responsibilities:

- assistant registry
- user identities
- domain management
- permission system
- conversation history
- A2A message logs
- company memory
- centralized RAG

The Brain allows assistants to discover who exists, who is responsible for what, and what information they can access.

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

Domains define organizational context:

- members
- responsibilities
- permissions
- accessible knowledge
- conversations

The same assistant can participate in multiple domains while maintaining separate contexts.

---

# Personal Harness

The assistant operates through a local harness.

The first interface is TUI-first, inspired by tools such as Claude Code and OpenCode.

The harness provides:

- terminal interface
- AI conversation
- local machine access
- file interaction
- command execution
- local context awareness

The objective is:

> One harness to interact with the assistant and accomplish work.

The first integration uses a DeepSeek-based harness approach. waïsp adds the collaboration layer around the harness instead of rebuilding an entire execution environment.

---

# A2A Collaboration

Assistants communicate through an existing A2A protocol.

Example flow:

```
Assistant A
     |
     | A2A
     |
Company Brain
     |
     |
Assistant B
```

The Brain handles:

- assistant discovery
- permission checks
- routing
- conversation storage
- context management

If the Brain is temporarily unavailable, assistants can operate using their latest cached permissions and synchronize later.

---

# Example Use Cases

## Intelligent Development Notification

A developer pushes new code.

The assistant analyzes the event and checks the company context.

It knows:

- who owns the backend
- who needs information
- who should not be notified

The assistant contacts only relevant assistants.

If the receiver prefers another language, the assistant adapts the communication automatically.

---

## Automatic Problem Resolution

A user encounters a server issue.

The harness provides local context.

The assistant identifies the problem and discovers the responsible person through the Company Brain.

It can ask permission from the user before contacting the responsible assistant.

The other assistant can answer with its current state, allowing collaboration without unnecessary human coordination.

---

# Deployment Model

## Phase 1 — LAN

The first deployment target is local enterprise networks.

Benefits:

- privacy
- self-hosting
- data ownership
- low latency

## Phase 2 — Online

Future versions can support internet-accessible Company Brains for distributed organizations.

---

# Technical Direction

Initial focus:

- Go-based core infrastructure
- TUI applications
- DeepSeek harness integration
- A2A communication
- self-hosted Company Brain
- modular architecture

The system should remain extensible:

- multiple harness providers
- additional AI models
- new tools
- enterprise integrations

---

# Long-Term Goal

waïsp aims to become the collaboration layer between humans and AI agents.

A future enterprise is not only a group of people using software. It is a network of people and assistants sharing context, coordinating work, and preserving organizational intelligence.
