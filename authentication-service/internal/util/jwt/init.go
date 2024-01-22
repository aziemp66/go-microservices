package jwt

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v4"

	errCommon "authentication/internal/http/error"
)

type JWTManager struct {
	AccessTokenKey []byte
}

func NewJWTManager(accessTokenKey string) *JWTManager {
	return &JWTManager{AccessTokenKey: []byte(accessTokenKey)}
}

func (j JWTManager) GenerateAuthToken(
	identifier string,
	name string,
	role string,
	duration time.Duration,
) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, AuthClaims{
		Name:       name,
		Identifier: identifier,
		Role:       role,
		RegisteredClaims: jwt.RegisteredClaims{
			// ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 365 * 10)),
		},
	})

	tokenString, err := token.SignedString(j.AccessTokenKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j JWTManager) VerifyAuthToken(tokenString string) (claim *AuthClaims, err error) {
	claims := &AuthClaims{}

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return j.AccessTokenKey, nil
	})
	if err != nil {
		err = errCommon.NewBadRequest("Invalid token : " + err.Error())
		return
	}

	if !tkn.Valid {
		err = errCommon.NewForbidden("You are not authorized for this acccess")
		return
	}

	return claims, nil
}
func (j JWTManager) GenerateTicketToken(ticketId string, duration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, TicketClaims{
		TicketId: ticketId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	})

	tokenString, err := token.SignedString(j.AccessTokenKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j JWTManager) VerifyTicketToken(tokenString string) (string, error) {
	claims := &TicketClaims{}

	stringAccessToken := string(j.AccessTokenKey)

	ticketAccessToken := []byte(stringAccessToken)

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return ticketAccessToken, nil
	})

	if err != nil {
		return "", err
	}

	if !tkn.Valid {
		return "", err
	}

	return claims.TicketId, nil
}
