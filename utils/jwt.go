package utils

import (
	"crypto/rsa"
	"errors"

	"github.com/dgrijalva/jwt-go"
)

// RS256 is RSA Signature with SHA-256.
const RS256 = "RS256"

// JWTHelper decodes or encodes JWT access token.
type JWTHelper interface {
	GenerateToken(claims jwt.Claims) (string, error)
	ParseClaims(token string, claims jwt.Claims) error
}

type jwtHelper struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
	parser     *jwt.Parser
}

// NewJWTHelper returns a new instance of JWTHelper.
func NewJWTHelper(publicKeyStr, privateKeyStr string) (JWTHelper, error) {
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKeyStr))
	if err != nil {
		return nil, err
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKeyStr))
	if err != nil {
		return nil, err
	}

	jwtHelper := &jwtHelper{
		privateKey: privateKey,
		publicKey:  publicKey,
		parser: &jwt.Parser{
			ValidMethods:         []string{RS256},
			SkipClaimsValidation: false,
		},
	}
	return jwtHelper, nil
}

func (h *jwtHelper) GenerateToken(claims jwt.Claims) (string, error) {
	if claims == nil {
		return "", errors.New("claims must not be nil")
	}

	tkn := jwt.NewWithClaims(jwt.GetSigningMethod(RS256), claims)
	str, err := tkn.SignedString(h.privateKey)
	if err != nil {
		return "", err
	}

	return str, nil
}

func (h *jwtHelper) ParseClaims(tokenStr string, claims jwt.Claims) error {
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (i interface{}, e error) {
		return h.publicKey, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return errors.New("invalid token")
	}
	return nil
}
