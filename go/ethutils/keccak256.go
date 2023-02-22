package ethutils

import (
	"fmt"
	"golang.org/x/crypto/sha3"
)

func Keccak256(data ...[]byte) string {
	keccak256 := sha3.NewLegacyKeccak256()
	for _, v := range data {
		keccak256.Write(v)
	}
	return fmt.Sprintf("%x", keccak256.Sum(nil))
}
