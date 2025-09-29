package helpers

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/SecretBin-io/go-secretbin/v2"
)

// ctxKeySecretBin is the context key used to store the SecretBin client in the command context.
type ctxKeySecretBin struct{}

// InitializeCLI initializes the CLI by creating a new SecretBin client and printing the server banner.
// This function is typically called in the PreRunE of all commands that need the client.
func InitializeCLI(cmd *cobra.Command, args []string) error {
	// Create a new SecretBin client using the configured endpoint
	sb, err := secretbin.New(viper.GetString("endpoint"))
	if err != nil {
		return err
	}

	// Print the SecretBin server banner if available
	PrintBanner(sb)

	// Store the SecretBin client in the command context for use in subcommands
	cmd.SetContext(context.WithValue(cmd.Context(), ctxKeySecretBin{}, sb))

	return nil
}

// GetSecretBinClient retrieves the SecretBin client from the command context.
// This is set in the PreRunE of the commands.
// It panics if the client is not found, which should not happen if the CLI is set up correctly.
func GetSecretBinClient(cmd *cobra.Command) secretbin.Client {
	sb, ok := cmd.Context().Value(ctxKeySecretBin{}).(secretbin.Client)
	if !ok {
		fmt.Fprintln(os.Stderr, "Error: SecretBin client not found in command context")
		os.Exit(1)
	}

	return sb
}
