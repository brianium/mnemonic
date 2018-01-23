package mnemonic

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brianium/mnemonic/entropy"
)

// Mnemonic represents a collection of human readable words
// used for HD wallet seed generation
type Mnemonic struct {
	Words []string
}

// New returns a new Mnemonic for the given entropy and language
func New(ent []byte, lang Language) *Mnemonic {
	const chunkSize = 11
	bits := entropy.CheckSummed(ent)
	length := len(bits)
	words := make([]string, length/11)
	for i := 0; i < length; i += chunkSize {
		stringVal := string(bits[i : chunkSize+i])
		intVal, err := strconv.ParseInt(stringVal, 2, 64)
		if err != nil {
			panic(fmt.Sprintf("Could not convert %s to word index", stringVal))
		}
		words[(chunkSize+i)/11-1] = GetWord(lang, intVal)
	}
	m := Mnemonic{words}
	return &m
}

// NewRandom returns a new Mnemonic with random entropy of the given length
// in bits
func NewRandom(length int, lang Language) *Mnemonic {
	ent, err := entropy.Random(length)
	if err != nil {
		panic(fmt.Sprintf("Error generating random entropy: %s", err))
	}
	return New(ent, lang)
}

// Sentence returns a Mnemonic's word collection as a space separated
// sentence
func (m *Mnemonic) Sentence() string {
	return strings.Join(m.Words, " ")
}

// GenerateSeed returns a seed used for wallet generation per
// BIP-0032 or similar method. The internal Words set
// of the Mnemonic will be used
func (m *Mnemonic) GenerateSeed(passphrase string) *Seed {
	return NewSeed(m.Sentence(), passphrase)
}
