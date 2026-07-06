package arguments

import (
	"flag"
	"fmt"
	"strings"
)

type Flags struct {
	Help bool
}

type Arguments struct {
	ComposeFilePath string
	Flags           Flags
}

// Parse commandline flags and arguments.
// Returns an Argument object containing all flags and the given file name
func Parse() Arguments {
	calledWithHelp := false
	flag.BoolFunc("h", "Show help", func(val string) error {
		printHelp()
		calledWithHelp = true
		return nil
	})

	flag.Parse()

	arguments := Arguments{
		ComposeFilePath: strings.Join(flag.Args(), " "),
		Flags: Flags{
			Help: calledWithHelp,
		},
	}
	return arguments
}

func printHelp() {
	fmt.Println("Help message WIP")
}
