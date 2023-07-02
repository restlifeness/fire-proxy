package security

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/scrypt"
)

const (
	SaltSize    = 16
	HashSize    = 64
	CostFactorN = 32768
	CostFactorR = 8
	CostFactorP = 1
	HashFormat  = "scrypt$%s$%s"
)

// generateSalt generates a new random salt.
func generateSalt() (string, error) {
	salt := make([]byte, SaltSize)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.RawStdEncoding.EncodeToString(salt), nil
}

// HashPassword hashes a password using the scrypt algorithm.
func HashPassword(password string) (string, error) {
	salt, err := generateSalt()
	if err != nil {
		return "", err
	}

	hash, err := scrypt.Key([]byte(password), []byte(salt), CostFactorN, CostFactorR, CostFactorP, HashSize)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(HashFormat, salt, base64.RawStdEncoding.EncodeToString(hash)), nil
}

// ComparePasswords compares a password with a hashed password.
func ComparePasswords(password string, hashedPassword string) (bool, error) {
	parts := strings.Split(hashedPassword, "$")
	if len(parts) != 3 {
		return false, fmt.Errorf("invalid hashed password format")
	}

	salt := parts[1]
	expectedHash := parts[2]

	hash, err := scrypt.Key([]byte(password), []byte(salt), CostFactorN, CostFactorR, CostFactorP, HashSize)
	if err != nil {
		return false, err
	}

	return base64.RawStdEncoding.EncodeToString(hash) == string(expectedHash), nil
}
