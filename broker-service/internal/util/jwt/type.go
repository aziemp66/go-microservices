package jwt

import jwt "github.com/golang-jwt/jwt/v4"

type AuthClaims struct {
	Identifier string
	Name       string
	Role       string
	jwt.RegisteredClaims
}
type TicketClaims struct {
	TicketId string `json:"ticket_id"`
	jwt.RegisteredClaims
}
