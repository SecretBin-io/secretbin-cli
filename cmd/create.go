package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Nihility-io/SecretBin-CLI/helpers"
	secretbin "github.com/Nihility-io/SecretBin-Go/v2"
)

// createCommand constructs the 'create' command for the CLI application.
func createCommand() *cobra.Command {
	var (
		// flagAttachments holds the list of file attachments to be added to the secret.
		// It is a slice of strings, allowing multiple files to be specified.
		flagAttachments []string

		// flagOptions holds the options for creating a secret, such as burn after reading,
		// password protection, and expiration time.
		flagOptions = secretbin.Options{}
	)
	cmd := &cobra.Command{
		Use:          "create [message]",
		Short:        "Create a new secret using SecretBin",
		SilenceUsage: true, // Don't print usage on error
		Args:         cobra.MinimumNArgs(1),
		PreRunE:      helpers.InitializeCLI,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Retrieve the SecretBin client from the command context.
			sb := helpers.GetSecretBinClient(cmd)

			// Create a secret with arguments joined into a single message.
			secret := secretbin.Secret{Message: strings.Join(args, " ")}

			// Append all provided files to the secret as attachments.
			for _, attachment := range flagAttachments {
				if err := secret.AddFileAttachment(attachment); err != nil {
					return err
				}
			}

			// Submit the secret with the specified options to SecretBin.
			// This will encrypt the secret and return a link to access it.
			link, err := sb.SubmitSecret(secret, flagOptions)
			if err != nil {
				return err
			}

			// Print the link to the created secret.
			fmt.Println(link)

			return nil
		},
	}

	cmd.Flags().
		StringSliceVarP(&flagAttachments, "attachment", "a", flagAttachments,
			"Attachment files to the secret (can be specified multiple times)")
	cmd.Flags().
		UintVarP(&flagOptions.BurnAfter, "burn-after", "b", flagOptions.BurnAfter,
			"Times the secret can be read before being deleted (0 means no burn after reading)")
	cmd.Flags().
		StringVarP(&flagOptions.Password, "password", "p", flagOptions.Password,
			"Additionally encrypt the secret with a password")
	cmd.Flags().
		StringVarP(&flagOptions.Expires, "expires", "x", flagOptions.Expires,
			"Expiration time for the secret (e.g. 1hr, 1d, 1w, 1m, 1y)")
	cmd.RegisterFlagCompletionFunc("expires",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			options := []string{}
			sb, err := secretbin.New(viper.GetString("endpoint"))
			if err != nil {
				return options, cobra.ShellCompDirectiveError
			}

			for k, v := range sb.Config().ExpiresSorted() {
				options = append(options, fmt.Sprintf("%s\t%s", k, v.String()))
			}

			return options, cobra.ShellCompDirectiveNoFileComp
		})

	return cmd
}
