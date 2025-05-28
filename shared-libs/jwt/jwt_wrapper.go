package jwt

import (
	"crypto/rsa"
	"fmt"
	"os"

	"github.com/IllhamHanafi/smart-traffic-routing-system/shared-libs/config"
	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

type JWTInterface interface {
	GenerateAndSignToken(claims jwt.Claims) (string, error)
}

func New(cfg config.JWT) (JWTInterface, error) {
	privateKeyFile, err := os.ReadFile(cfg.RSAPrivateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key file: %w", err)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyFile)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	publicKeyFile, err := os.ReadFile(cfg.RSAPublicKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read public key file: %w", err)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyFile)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	return &service{
		privateKey: privateKey,
		publicKey:  publicKey,
	}, nil
}

func (s *service) GenerateAndSignToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err := token.SignedString(s.privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}
	return signedToken, nil
}

// func (s *service) ValidateToken(token string) (jwt.MapClaims, error) {

// }
