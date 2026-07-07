package main

import (
	// "fmt"
	"github.com/lauri-lyytikainen/composemap/arguments"
	"github.com/lauri-lyytikainen/composemap/parser"
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

	// fmt.Println(runSettings.ComposeFilePath)
	parser.ParseComposeFile(runSettings.ComposeFilePath)
}
