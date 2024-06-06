package token

import (
	"testing"
	"time"

	"github.com/Aashish-32/bank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

}

func (m *PasetoMaker) TestVerifyToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute
	Issued_at := time.Now()
	expired_at := time.Now().Add(duration)

	token, err := m.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.Equal(t, username, payload.Username)

	require.NotZero(t, payload.ID)
	require.WithinDuration(t, Issued_at, payload.Issued_at, time.Second)
	require.WithinDuration(t, expired_at, payload.Expiry, time.Second)

}
