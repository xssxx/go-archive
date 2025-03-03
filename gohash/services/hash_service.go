package services

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"testptr/interfaces"
)

type Sha256Hasher struct{}

func (s *Sha256Hasher) HashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

type Sha512Hasher struct{}

func (s *Sha512Hasher) HashPassword(password string) string {
	hash := sha512.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

func NewSha256Hasher() interfaces.PasswordHasher {
	return &Sha256Hasher{}
}

func NewSha512Hasher() interfaces.PasswordHasher {
	return &Sha512Hasher{}
}
