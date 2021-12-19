package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"fmt"

	"github.com/WebchemistGenn/WebchemistCoin/utils"
)

const (
	signature     string = "2d6020f9ae4142dd2b9165bc187ca55b5364195fd7c0b27c66abd62d7fd305593ade6bf7ca6448b37d0b097815c5b722567c69dc8ddf48bbbe8102031e60fb43⏎"
	privateKey    string = "30770201010420335bdfe5a7848197446a6d3c7f649a28e1fc8dd116f311a3bd36cf7b45df05cca00a06082a8648ce3d030107a14403420004410ea51ad2ae9f7db808dc260902d3e6a9d700f6ea1124c2793e72493a98a6c8dad837c5a79f7a853a9d96fc6f2727478e3dc380a7d060fbbdfa58c6fd03a022"
	hashedMessage string = "1c5863cd55b5a4413fd59f054af57ba3c75c0698b3851d70f99b8de2d5c7338f"
)

func Start() {
	privateByte, err := hex.DecodeString(privateKey)
	utils.HandleErr(err)

	restoredKey, err := x509.ParseECPrivateKey(privateByte)
	utils.HandleErr(err)

	fmt.Println(restoredKey)
}
