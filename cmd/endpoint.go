package cmd

import (
	"github.com/spf13/cobra"

	"github.com/secretbin-io/secretbin-cli/helpers"
)

// endpointCommand constructs the 'set-endpoint' command for the CLI application.
func endpointCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "set-endpoint <endpoint>",
		Short:        "Set where SecretBin is hosted",
		SilenceUsage: true, // Don't print usage on error
		Args:         cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Set the endpoint for SecretBin using the provided argument.
			return helpers.SetEndpoint(args[0])
		},
	}

	return cmd
}
