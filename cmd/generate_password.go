package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	secretbin "github.com/Nihility-io/SecretBin-Go/v2"
)

// generatePasswordCommand constructs the 'generate-password' command for the CLI application.
func generatePasswordCommand() *cobra.Command {
	var (
		flagLength      uint = 16
		flagNoUppercase bool = false
		flagNoLowercase bool = false
		flagNoDigits    bool = false
		flagNoSymbols   bool = false
	)

	cmd := &cobra.Command{
		Use:   "generate-password",
		Short: "Generates a secure password",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			pw, err := secretbin.GeneratePassword(secretbin.PasswordOptions{
				Length:    int(flagLength),
				Uppercase: !flagNoUppercase,
				Lowercase: !flagNoLowercase,
				Digits:    !flagNoDigits,
				Symbols:   !flagNoSymbols,
			})

			if err != nil {
				return err
			}

			fmt.Fprintln(os.Stdout, pw)

			return nil
		},
	}

	cmd.Flags().
		UintVarP(&flagLength, "length", "l", flagLength,
			"Length of the generated password")
	cmd.Flags().
		BoolVar(&flagNoUppercase, "no-uppercase", flagNoUppercase,
			"Exclude uppercase letters from the generated password")
	cmd.Flags().
		BoolVar(&flagNoLowercase, "no-lowercase", flagNoLowercase,
			"Exclude lowercase letters from the generated password")
	cmd.Flags().
		BoolVar(&flagNoDigits, "no-digits", flagNoDigits,
			"Exclude digits from the generated password")
	cmd.Flags().
		BoolVar(&flagNoSymbols, "no-symbols", flagNoSymbols,
			"Exclude symbols from the generated password")

	return cmd
}
