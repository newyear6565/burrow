package crypto

import (
	"crypto/sha256"

	"github.com/hyperledger/burrow/crypto/gmssl/sm3"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

func Keccak256(data []byte) []byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(data)
	return hash.Sum(nil)
}

func SHA256(data []byte) []byte {
	hash := sha256.New()
	hash.Write(data)
	return hash.Sum(nil)
}

func RIPEMD160(data []byte) []byte {
	hash := ripemd160.New()
	hash.Write(data)
	return hash.Sum(nil)
}

func SM3(data []byte) []byte {
	sm3hash := sm3.New()
	sm3hash.Write(data)
	return sm3hash.Sum(nil)
}
