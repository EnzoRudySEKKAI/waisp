package a2a

import "github.com/EnzoRudySEKKAI/waisp/internal/domain"

// Messenger defines the communication boundary between assistants.
// Implementations can later use the selected A2A protocol transport.
type Messenger interface {
	Send(message domain.Message) error
}

// LocalMessenger is the first in-memory implementation used by the core.
type LocalMessenger struct {
	Messages []domain.Message
}

func (m *LocalMessenger) Send(message domain.Message) error {
	m.Messages = append(m.Messages, message)
	return nil
}
