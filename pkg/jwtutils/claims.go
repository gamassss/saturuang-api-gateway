package jwtutils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type CustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (c CustomClaims) Valid() error {
	if c.RegisteredClaims.ExpiresAt != nil && time.Now().After(c.RegisteredClaims.ExpiresAt.Time) {
		return jwt.ErrTokenExpired
	}
	return nil
}
