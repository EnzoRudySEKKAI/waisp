package brain

import (
	"github.com/EnzoRudySEKKAI/waisp/internal/domain"
)

// Brain is the company-level coordination layer.
// Storage, RAG and permissions will be plugged behind interfaces later.
type Brain struct {
	Users       map[string]domain.User
	Assistants  map[string]domain.Assistant
	Domains     map[string]domain.Domain
}

func New() *Brain {
	return &Brain{
		Users:      make(map[string]domain.User),
		Assistants: make(map[string]domain.Assistant),
		Domains:    make(map[string]domain.Domain),
	}
}
