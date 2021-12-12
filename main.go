package main

import (
	"github.com/WebchemistGenn/WebchemistCoin/blockchain"
	"github.com/WebchemistGenn/WebchemistCoin/cli"
	"github.com/WebchemistGenn/WebchemistCoin/database"
)

func main() {
	defer database.Close()
	blockchain.Blockchain()
	cli.Start()
}
