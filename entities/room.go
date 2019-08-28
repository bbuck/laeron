package entities

import (
	"github.com/google/uuid"
)

// Room represents an entity that players are intended to be present in.
type Room struct {
	ID          uuid.UUID `json:"ID"`
	Name        string    `json:"Name"`
	Description string    `json:"Description"`
	AreaID      uuid.UUID `json:"UUID"`
}
