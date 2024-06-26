package token

import (
	"time"

	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var errExpiredToken = errors.New("token has expired")

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

func (p *Payload) Valid() error {
	if time.Now().After(p.Expiry) {
		return errExpiredToken
	}
	return nil
}

///BELOW METHODS ARE WRITTEN JUST TO IMPLEMENT JWT.CLAIMS INTERFACE

func (p *Payload) GetAudience() (jwt.ClaimStrings, error) {
	if p == nil {
		return nil, errors.New("nil pointer receiver")
	}
	return []string{p.Username}, nil
}

func (p *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	if p == nil {
		return jwt.NewNumericDate(time.Time{}), errors.New("nil pointer receiver")
	}
	return jwt.NewNumericDate(p.Expiry), nil
}

func (p *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	if p == nil {
		return jwt.NewNumericDate(time.Time{}), errors.New("nil pointer receiver")
	}
	return jwt.NewNumericDate(p.Issued_at), nil

}

func (p *Payload) GetIssuer() (string, error) {
	if p == nil {
		return "", errors.New("nil pointer receiver")
	}
	return "your_issuer", nil

}
func (p *Payload) GetSubject() (string, error) {
	if p == nil {
		return "", errors.New("nil pointer receiver")
	}
	return p.ID.String(), nil

}
func (p *Payload) GetNotBefore() (*jwt.NumericDate, error) {
	if p == nil {
		return jwt.NewNumericDate(time.Time{}), errors.New("nil pointer receiver")
	}
	return jwt.NewNumericDate(time.Time{}), nil

}
