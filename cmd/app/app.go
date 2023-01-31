package app

import (
	"context"
	"github.com/spf13/cobra"
)

const (
	lDescription = `
This CLI helps to automate the action of creating tickets when you need to do it in a recurring matter.
Just create instruction in YAML, pass it to the CLI, provide your credential, project, and sit back.
`
	examples = `
 # Execute dry-run
 jiratron --config=config.yaml --access-token=token --dry-run

 # Actually create epics/tickets
 jiratron --config=config.yaml --access-token=token
`
)

func New(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:     "jiratron",
		Example: examples,
		Short:   "Create an epic in Jira from yaml file with the descriptions, links, story points",
		Long:    lDescription,
		RunE:    Run,
	}
}

func Run(cmd *cobra.Command, args []string) error {
	return nil
}
