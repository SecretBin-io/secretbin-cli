package cmd

import (
	"github.com/Nihility-io/SecretBin-CLI/helpers"
	"github.com/spf13/cobra"
)

var (
	setEndpointCmd = &cobra.Command{
		Use:          "set-endpoint <endpoint>",
		Short:        "Set where SecretBin is hosted",
		SilenceUsage: true, // Don't print usage on error
		Args:         cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Set the endpoint for SecretBin using the provided argument.
			return helpers.SetEndpoint(args[0])
		},
	}
)
