package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"

	"github.com/SecretBin-io/secretbin-cli/cmd"
	"github.com/SecretBin-io/secretbin-cli/helpers"
)

var (
	NAME    = "secretbin" // Name of the CLI application (set by the build system)
	VERSION = "0.0.0"     // Version of the CLI application (set by the build system)
)

func init() {
	// Get the user's configuration directory and create a subdirectory for the CLI.
	configDir, _ := os.UserConfigDir()
	configDir = filepath.Join(configDir, NAME)
	os.MkdirAll(configDir, 0o755)

	// Set the configuration file name and type for viper.
	viper.SetConfigType("toml")
	viper.AddConfigPath(configDir)
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	// Require the user to the set the SecretBin endpoint if it is not already configured.
	if viper.GetString("endpoint") == "" {
		for {
			endpoint := ""
			print("Type the SecretBin endpoint (e.g. https://secretbin.example.com): ")
			_, err := fmt.Scanln(&endpoint)
			if err != nil {
				continue
			}

			if err := helpers.SetEndpoint(endpoint); err != nil {
				fmt.Fprintf(os.Stderr, "Invalid endpoint: %s\n", err)

				continue
			}

			break
		}
	}
}

func main() {
	// Initialize the CLI by setting up the root command and its subcommands.
	cmd.Execute(NAME, VERSION)
}
