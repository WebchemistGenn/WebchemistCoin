package main

import (
	"github.com/WebchemistGenn/WebchemistCoin/explorer"
	"github.com/WebchemistGenn/WebchemistCoin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
