package utils

import (
	_ "crypto/rsa"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key") // replace with your own secret key

// CreateToken generates a JWT token with the provided payload and TTL, signed with the given private key
func CreateToken(ttl time.Duration, payload interface{}, privateKey string) (string, error) {
	keyBytes, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", fmt.Errorf("could not decode key: %w", err)
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(keyBytes)
	if err != nil {
		return "", fmt.Errorf("could not parse private key: %w", err)
	}

	now := time.Now().UTC()

	claims := jwt.MapClaims{
		"sub": payload,
		"exp": now.Add(ttl).Unix(),
		"iat": now.Unix(),
		"nbf": now.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err := token.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("could not sign token: %w", err)
	}

	return signedToken, nil
}

// ValidateToken parses and validates a JWT token with the provided public key
func ValidateToken(token string, publicKey string) (interface{}, error) {
	keyBytes, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, fmt.Errorf("could not decode key: %w", err)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(keyBytes)
	if err != nil {
		return nil, fmt.Errorf("could not parse public key: %w", err)
	}

	parser := jwt.Parser{
		ValidMethods: []string{jwt.SigningMethodRS256.Name},
	}

	parsedToken, err := parser.ParseWithClaims(token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("could not parse token: %w", err)
	}

	if !parsedToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims["sub"], nil
}


func GenerateToken(userID uint) (string, error) {
	// Set the expiration time of the token to 24 hours from now
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create the JWT claims, which includes the user ID and expiration time
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     expirationTime.Unix(),
	}

	// Create the JWT token with the claims and sign it with the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("error creating token: %w", err)
	}

	return tokenString, nil
}