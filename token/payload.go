package token

import (
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Issued_at time.Time `json:"created_at"`
	Expiry    time.Time `json:"expiry"`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		ID:        id,
		Username:  username,
		Issued_at: time.Now(),
		Expiry:    time.Now().Add(duration),
	}
	return payload, nil

}
