package main

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
)

// SeedRand generates a Rand object.
func SeedRand() *rand.Rand {
	var b [8]byte
	_, err := crand.Read(b[:])
	if err != nil {
		panic("cannot seed with cryptographic random number generator")
	}
	r := rand.New(rand.NewSource(int64(binary.LittleEndian.Uint64(b[:]))))
	return r
}

func main() {
	r := SeedRand()
	num := r.Int()
	fmt.Println(num)
}
