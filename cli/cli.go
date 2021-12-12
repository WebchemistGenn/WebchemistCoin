package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/WebchemistGenn/WebchemistCoin/explorer"
	"github.com/WebchemistGenn/WebchemistCoin/rest"
)

func usage() {
	fmt.Printf("Welcome to Webchemist Coin\n\n")
	fmt.Printf("Please use the following commands:\n\n")
	fmt.Printf("-port=4000: 		Set the PORT of the SERVER\n")
	fmt.Printf("-mode=rest: 		Choose between 'html' and 'rest'\n\n")
	os.Exit(0)
}
func Start() {
	if len(os.Args) < 2 {
		usage()
	}

	port := flag.Int("port", 4000, "Set the PORT of the SERVER")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}
}
