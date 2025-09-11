package cmd

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"

	"github.com/Nihility-io/SecretBin-CLI/helpers"
)

// infoCommand constructs the 'info' command for the CLI application.
func infoCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "info",
		Short:        "Print information about the SecretBin server",
		SilenceUsage: true, // Don't print usage on error
		Args:         cobra.NoArgs,
		PreRunE:      helpers.InitializeCLI,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Retrieve the SecretBin client from the command context.
			sb := helpers.GetSecretBinClient(cmd)

			// Print information about the SecretBin client and server in a table format.
			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			t.SetStyle(table.StyleRounded)

			t.AppendRow(table.Row{"Information", "Information"}, table.RowConfig{AutoMerge: true})
			t.AppendSeparator()
			t.AppendRows([]table.Row{
				{"Endpoint", sb.Config().Endpoint},
				{"Client Version", fmt.Sprintf("%s %s", cmd.Root().Use, cmd.Root().Version)},
				{"Server Version", fmt.Sprintf("%s %s", sb.Config().Name, sb.Config().Version)},
			})

			t.AppendSeparator()

			// Print the available expiration times in a table format.
			t.AppendRow(table.Row{"Expiration Times", "Expiration Times"}, table.RowConfig{AutoMerge: true})
			t.AppendSeparator()
			for k, v := range sb.Config().ExpiresSorted() {
				if k == sb.Config().DefaultExpires {
					k = text.FgGreen.Sprint(k + " (default)")
				}
				t.AppendRow(table.Row{k, v.String()})
			}

			t.Render()

			return nil
		},
	}

	return cmd
}
