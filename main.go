package main

import (
	"fmt"
	"github.com/lauri-lyytikainen/composemap/arguments"
)

func main() {
	args := arguments.Parse()
	fmt.Println(args.Flags.Help)
	fmt.Println(args.ComposeFilePath)
	// loadComposeFile()
	// printOutput()
}
