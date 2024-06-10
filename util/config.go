package util

import "time"

type Config struct {
	TokenSymmetricKey   string
	AccessTokenDuration time.Duration
}
