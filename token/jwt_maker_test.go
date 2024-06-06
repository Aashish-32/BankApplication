package token

import (
	"fmt"
	"testing"
	"time"

	"github.com/Aashish-32/bank/util"
	"github.com/stretchr/testify/require"
)

func TestJwtmaker(t *testing.T) {

	maker, err := NewJWTMaker(util.RandomString(22))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute
	Issued_at := time.Now()
	expired_at := Issued_at.Add(duration)

	token, err := maker.CreateToken(username, duration)
	fmt.Println(token)
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

func TestExpiredJWTToken(t *testing.T) {

	maker, err := NewJWTMaker(util.RandomString(22))
	require.NoError(t, err)

	token, err := maker.CreateToken(util.RandomOwner(), -time.Minute)
	fmt.Println(token)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	fmt.Println(err)
	require.Error(t, err)
	require.EqualError(t, err, fmt.Errorf("token has invalid claims: token is expired").Error())

	require.Nil(t, payload)

}
