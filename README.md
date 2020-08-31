我是光年实验室高级招聘经理。
我在github上访问了你的开源项目，你的代码超赞。你最近有没有在看工作机会，我们在招软件开发工程师，拉钩和BOSS等招聘网站也发布了相关岗位，有公司和职位的详细信息。
我们公司在杭州，业务主要做流量增长，是很多大型互联网公司的流量顾问。公司弹性工作制，福利齐全，发展潜力大，良好的办公环境和学习氛围。
公司官网是http://www.gnlab.com,公司地址是杭州市西湖区古墩路紫金广场B座，若你感兴趣，欢迎与我联系，
电话是0571-88839161，手机号：18668131388，微信号：echo 'bGhsaGxoMTEyNAo='|base64 -D ,静待佳音。如有打扰，还请见谅，祝生活愉快工作顺利。

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
