package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Command creates the root command for the CLI application with the given name and version.
func Command(name string, version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:          name,
		Version:      version,
		Short:        "SecretBin Command Line Interface",
		SilenceUsage: true, // Don't print usage on error
		Args:         cobra.NoArgs,
	}

	// Add subcommands to the root command.
	cmd.AddCommand(createCommand())
	cmd.AddCommand(endpointCommand())
	cmd.AddCommand(generatePasswordCommand())
	cmd.AddCommand(infoCommand())

	// Persistent flags are flags that will be available to all subcommands.
	cmd.PersistentFlags().
		Bool("hide-banner", false, "Do not show the SecretBin banner message. This is useful for scripting.")

	// Bind the persistent flags to viper for configuration management.
	viper.BindPFlags(cmd.PersistentFlags())

	return cmd
}

// Execute runs the root command and sets the name and version of the CLI.
// It will exit with status code 1 if there is an error during execution.
func Execute(name string, version string) {
	if err := Command(name, version).Execute(); err != nil {
		os.Exit(1)
	}
}
