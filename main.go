package main

import (
	"fmt"
	"github.com/lauri-lyytikainen/composemap/arguments"
)

func main() {
	args, usageFunc, err := arguments.Parse()
	if err != nil {
		panic(err)
	}
	runSettings := arguments.HandleArgs(args, usageFunc)

	if runSettings.CanReturnEarly {
		return
	}

	fmt.Println(runSettings.ComposeFilePath)
	// loadComposeFile()
	// printOutput()
}
