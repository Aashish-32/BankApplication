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

func TestVerifyToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute
	Issued_at := time.Now()
	expired_at := time.Now().Add(duration)

	token, err := maker.CreateToken(username, duration)
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
func TestExpiredPasteoToken(t *testing.T) {

	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	token, err := maker.CreateToken(util.RandomOwner(), -time.Minute)

	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, errExpiredToken.Error())

	require.Nil(t, payload)

}
