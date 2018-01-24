# Mnemonic
[![Build Status](https://travis-ci.org/brianium/mnemonic.svg?branch=master)](https://travis-ci.org/brianium/mnemonic)
[![Go Report Card](https://goreportcard.com/badge/github.com/brianium/mnemonic)](https://goreportcard.com/report/github.com/brianium/mnemonic)
[![GoDoc](https://godoc.org/github.com/brianium/mnemonic?status.svg)](https://godoc.org/github.com/brianium/mnemonic)

A BIP 39 implementation in Go.

Features:

* Generating human readable sentences for seed generation - a la [BIP 32](https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki)
* All languages mentioned in the [proposal](https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki) supported.
* 128 bit (12 words) through 256 bit (24 words) entropy.

## [`mnemonic`](https://godoc.org/github.com/brianium/mnemonic) package

* Generates human readable sentences and the seeds derived from them.
* Supports all languages mentioned in the [BIP 39 proposal](https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki).
* Supports ideogrpahic spaces for Japanese language.

Example:

```go
package main

import (
    "fmt"
    "github.com/brianium/mnemonic"
)

func main() {
    // generate a random Mnemonic in English with 256 bits of entropy
    m, _ := mnemonic.NewRandom(256, mnemonic.English)

    // print the Mnemonic as a sentence
    fmt.Println(m.Sentence())

    // inspect underlying words
    fmt.Println(m.Words)

    // generate a seed from the Mnemonic
    seed := m.GenerateSeed("passphrase")

    // print the seed as a hex encoded string
    fmt.Println(seed)
}
```

## [`entropy`](https://godoc.org/github.com/brianium/mnemonic/entropy) package

* Supports generating random entropy in the range of 128-256 bits
* Supports generating entropy from a hex string

Example:

```go
package main

import (
    "fmt"
    "github.com/brianium/mnemonic"
    "github.com/brianium/mnemonic/entropy"
)

func main() {
    // generate some entropy from a hex string
    ent, _ := entropy.FromHex("8197a4a47f0425faeaa69deebc05ca29c0a5b5cc76ceacc0")
    
    // generate a Mnemonic in Japanese with the generated entropy
    jp, _ := mnemonic.New(ent, mnemonic.Japanese)

    // print the Mnemonic as a sentence
    fmt.Println(jp.Sentence())

    // generate some random 256 bit entropy
    rnd, _ := entropy.Random(256)
    
    // generate a Mnemonic in Spanish with the generated entropy
    sp, _ := mnemonic.New(rnd, mnemonic.Spanish)

    // print the Mnemonic as a sentence
    fmt.Println(sp.Sentence())
}
```

# Installation

To install Mnemonic, use `go get`:

    go get github.com/brianium/mnemonic

This will then make the following packages available to you:

    github.com/brianium/mnemonic
    github.com/brianium/mnemonic/entropy

Import the `mnemonic` package into your code using this template:

```go
package yours

import (
  "github.com/brianium/mnemonic"
)

func MnemonicJam(passphrase string) {

  m := mnemonic.NewRandom(passphrase)

}
```

# Contributing

Please feel free to submit issues, fork the repository and send pull requests!

When submitting an issue, we ask that you please include a complete test function that demonstrates the issue.
