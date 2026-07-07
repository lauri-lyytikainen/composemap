package parser

import (
	"context"
	"fmt"
	"github.com/compose-spec/compose-go/v2/cli"
	"log"
)

func ParseComposeFile(path string) {

	ctx := context.Background()
	options, err := cli.NewProjectOptions(
		[]string{path},
		// cli.WithOsEnv,
		// cli.WithDotEnv,
	)

	if err != nil {
		log.Fatal(err)
	}

	project, err := options.LoadProject(ctx)
	if err != nil {
		log.Fatal(err)
	}

	projectYAML, err := project.MarshalYAML()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(projectYAML))
}
