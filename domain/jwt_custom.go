package domain

import (
	"github.com/golang-jwt/jwt/v4"
)

// JwtCustomClaims represents custom JWT claims.
type JwtCustomClaims struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	jwt.StandardClaims
}

// JwtCustomRefreshClaims represents custom JWT refresh claims.
type JwtCustomRefreshClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}
