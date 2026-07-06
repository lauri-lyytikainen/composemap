package arguments

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type Flags struct {
	Help  bool
	Input string
}

type Arguments struct {
	ComposeFilePath   string
	Flags             Flags
	UsageMessageShown bool
}

type RunSettings struct {
	ComposeFilePath string
	CanReturnEarly  bool
}

// Parse commandline flags and arguments.
// Returns an Argument object containing all flags and the given file name
func Parse() (Arguments, func(), error) {
	return ParseArgs(os.Args[1:], os.Stdout)
}

func ParseArgs(args []string, w io.Writer) (Arguments, func(), error) {
	fs := flag.NewFlagSet("composemap", flag.ContinueOnError)
	fs.SetOutput(w)

	usageShown := false

	helpFlag := fs.Bool("h", false, "Show help message")
	inputFlag := fs.String("i", "", "Specify Docker Compose file")

	fs.Usage = func() {
		fmt.Fprintf(w, "Usage: composemap [flags] <compose-file> or use the -i flag\n\r")
		fs.PrintDefaults()
	}

	err := fs.Parse(args)
	if err != nil {
		usageShown = true
	}

	filePath := *inputFlag
	if filePath == "" {
		filePath = strings.Join(fs.Args(), " ")
	}

	arguments := Arguments{
		ComposeFilePath: filePath,
		Flags: Flags{
			Help:  *helpFlag,
			Input: *inputFlag,
		},
		UsageMessageShown: usageShown,
	}
	return arguments, fs.Usage, nil
}

// Handle passed arguments and evaluate given command flags
func HandleArgs(args Arguments, usageFunc func()) RunSettings {

	if args.Flags.Help && !args.UsageMessageShown {
		usageFunc()
	}

	return RunSettings{
		ComposeFilePath: args.ComposeFilePath,
		CanReturnEarly:  args.UsageMessageShown || args.Flags.Help,
	}
}
