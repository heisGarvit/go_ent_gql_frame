package random

import (
	"hash/maphash"
	"math/rand"
)

var letterBytes = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func Gen(n int) string {
	r := rand.New(rand.NewSource(int64(new(maphash.Hash).Sum64())))
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[r.Intn(len(letterBytes))]
	}
	return string(b)
}
