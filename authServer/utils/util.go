package utils

import (
	"encoding/base64"
	"encoding/pem"
	"log"
	"math/rand"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	cost           = 14
	privateKeyPath = "keys/id_rsa"
	publicKeyPath  = "keys/id_rsa.pub"
)

var (
	seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
)

const (
	Charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandString(length int) string {
	return StringWithCharset(length, Charset)
}

// Used for hashing  password using base64 encoding
func PasswordEncoder(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", nil
	}
	return string(hash), nil
}

// Comparing password with the hash
func ComparePassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func ReadPrivateKey() (string, error) {
	keyData, err := os.ReadFile(privateKeyPath)
	if err != nil {
		log.Printf("ERROR: fail get idrsa, %s", err.Error())
		return "", err
	}

	keyBlock, _ := pem.Decode(keyData)
	if keyBlock == nil {
		log.Printf("ERROR: fail get idrsa, invalid key")
		return "", err
	}
	// encode base64 key data
	return base64.StdEncoding.EncodeToString(keyBlock.Bytes), nil
}
