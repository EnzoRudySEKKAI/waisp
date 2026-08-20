# A2A Communication

## Goal

Assistants communicate using an existing Agent-to-Agent protocol.

## Flow

```
Assistant A
 |
A2A Protocol
 |
Company Brain
 |
Assistant B
```

## Brain Responsibilities

- discover assistants
- verify permissions
- route messages
- store history
- preserve context

If unavailable, assistants can use cached permissions temporarily.
