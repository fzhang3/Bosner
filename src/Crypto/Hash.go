package Crypto

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func ComputeHashSha256(bytes []byte) string {
	return fmt.Sprintf("%x", sha256.Sum256(bytes))
}
