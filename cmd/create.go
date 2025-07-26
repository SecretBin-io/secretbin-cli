package cmd

import (
	"fmt"
	"strings"

	"github.com/Nihility-io/SecretBin-CLI/helpers"
	secretbin "github.com/Nihility-io/SecretBin-Go/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// flagCreateAttachments holds the list of file attachments to be added to the secret.
	// It is a slice of strings, allowing multiple files to be specified.
	flagCreateAttachments []string

	// flagCreateOptions holds the options for creating a secret, such as burn after reading,
	// password protection, and expiration time.
	flagCreateOptions = secretbin.Options{}
)

var (
	createCmd = &cobra.Command{
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
			for _, attachment := range flagCreateAttachments {
				if err := secret.AddFileAttachment(attachment); err != nil {
					return err
				}
			}

			// Submit the secret with the specified options to SecretBin.
			// This will encrypt the secret and return a link to access it.
			link, err := sb.SubmitSecret(secret, flagCreateOptions)
			if err != nil {
				return err
			}

			// Print the link to the created secret.
			fmt.Println(link)

			return nil
		},
	}
)

func init() {
	createCmd.Flags().StringSliceVarP(&flagCreateAttachments, "attachment", "a", flagCreateAttachments, "Attachment files to the secret (can be specified multiple times)")
	createCmd.Flags().UintVarP(&flagCreateOptions.BurnAfter, "burn-after", "b", flagCreateOptions.BurnAfter, "Times the secret can be read before being deleted (0 means no burn after reading)")
	createCmd.Flags().StringVarP(&flagCreateOptions.Password, "password", "p", flagCreateOptions.Password, "Additionally encrypt the secret with a password")
	createCmd.Flags().StringVarP(&flagCreateOptions.Expires, "expires", "x", flagCreateOptions.Expires, "Expiration time for the secret (e.g. 1hr, 1d, 1w, 1m, 1y)")
	createCmd.RegisterFlagCompletionFunc("expires", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
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
}
