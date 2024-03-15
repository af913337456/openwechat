// Package random is a small library for generating random strings
package random

import (
	cryptrand "crypto/rand"
	"encoding/hex"
	mathrand "math/rand"
	"time"
	"unsafe"
)

const (
	lowerCase = "abcdefghijklmnopqrstuvwxyz"
	upperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	number    = "1234567890"
	symbols   = `!@#$%^&*()_+{}|:<>?-=[]\;',./`
)

const (
	letterIdxBits = 6
	letterIdxMask = 2<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

// Option let you select additional letters to be added in generated random string
type Option struct {
	IncludeNumber    bool
	IncludeUpperCase bool
	IncludeSymbols   bool
}

// RandStr generate random string with length n, only use lower case latin alphabet
func RandStr(n int) string {
	return generateStr(n, Option{})
}

// RandStrWithOpt same as RandStr, but with option
func RandStrWithOpt(n int, opt Option) string {
	return generateStr(n, opt)
}

// RandHexStr generate random hexadecimals string with length n
func RandHexStr(n int) string {
	bytes := make([]byte, n)
	cryptrand.Read(bytes)

	return hex.EncodeToString(bytes)
}

func generateStr(n int, opt Option) string {
	letters := lowerCase
	if opt.IncludeNumber {
		letters += number
	}

	if opt.IncludeUpperCase {
		letters += upperCase
	}

	if opt.IncludeSymbols {
		letters += symbols
	}

	src := mathrand.NewSource(time.Now().UnixNano())

	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letters) {
			b[i] = letters[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}
