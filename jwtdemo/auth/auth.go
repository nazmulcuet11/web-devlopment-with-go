package auth

import (
	"crypto/rsa"
	"errors"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	privateKeyPath = "auth/keys/app.rsa"
	publicKeyPath  = "auth/keys/app.rsa.pub"
)

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

func init() {
	privateKeyBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatal("Failed to read priavte key from path: ", privateKeyPath)
		return
	}
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		log.Fatal("Failed to parse private key, error: ", err)
		return
	}

	publicKeyBytes, err := os.ReadFile(publicKeyPath)
	if err != nil {
		log.Fatal("Failed to read public key from path: ", publicKeyPath)
		return
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		log.Fatal("Failed to parse public key, error: ", err)
		return
	}
}

func GenerateToken(c *Credentials) (string, error) {
	expirationTime := time.Now().Add(time.Minute * 5).Unix()
	claims := &Claims{
		Username: c.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
			IssuedAt:  time.Now().Unix(),
			Issuer:    "my-awsome-company",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(privateKey)
	return tokenString, err
}

func VerifyToken(tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return errors.New("invalid token")
	}
	return nil
}
