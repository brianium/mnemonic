package mnemonic

import (
	"crypto/sha512"
	"encoding/hex"

	"golang.org/x/crypto/pbkdf2"
)

// Seed represents a binary seed used for HD wallet generation
type Seed struct {
	Bytes []byte
}

// NewSeed creae a new Seed with the given sentence and passphrase
func NewSeed(sentence string, passphrase string) *Seed {
	s := pbkdf2.Key([]byte(sentence), []byte("mnemonic"+passphrase), 2048, 64, sha512.New)
	seed := Seed{s}
	return &seed
}

// ToHex returns the seed bytes as a hex encoded string
func (s *Seed) ToHex() string {
	src := s.Bytes
	seedHex := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(seedHex, s.Bytes)
	return string(seedHex)
}

func (s *Seed) String() string {
	return s.ToHex()
}
