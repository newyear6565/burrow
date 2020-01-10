package crypto

import (
	"fmt"
	"github.com/hyperledger/burrow/crypto/gmssl"
)

func Sm2p256v1() {
	versions := gmssl.GetVersions()
	fmt.Println("GmSSL Versions:")
	for _, version := range versions {
		fmt.Println(" " + version)
	}
	fmt.Println("")
}
