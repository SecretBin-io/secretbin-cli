package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Short:        "SecretBin Command Line Interface",
		SilenceUsage: true, // Don't print usage on error
		Args:         cobra.NoArgs,
	}
)

// init initializes the root command and its subcommands.
func init() {
	// Add subcommands to the root command.
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(setEndpointCmd)

	// Persistent flags are flags that will be available to all subcommands.
	rootCmd.PersistentFlags().Bool("hide-banner", false, "Do not show the SecretBin banner message. This is useful for scripting.")

	// Bind the persistent flags to viper for configuration management.
	viper.BindPFlags(rootCmd.PersistentFlags())
}

// Execute runs the root command and sets the name and version of the CLI.
// It will exit with status code 1 if there is an error during execution.
func Execute(name string, version string) {
	rootCmd.Use = name
	rootCmd.Version = version
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
