package helpers

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/viper"

	secretbin "github.com/Nihility-io/SecretBin-Go/v2"
)

// PrintBanner prints the banner message of the SecretBin server if enabled in the configuration.
func PrintBanner(sb secretbin.Client) {
	// Check if the server displays a banner and if the hide-banner flag is not set
	if banner := sb.Config().Banner; banner != nil && !viper.GetBool("hide-banner") {
		// Determine the color based on the banner type
		bannerColor := map[string]text.Color{
			"info":    text.FgBlue,
			"warning": text.FgYellow,
			"error":   text.FgRed,
		}[banner.Type]

		// Create a new table writer for formatting the banner output
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)

		// Set the table style and color
		t.SetStyle(table.StyleRounded)
		t.Style().Color.Border = text.Colors{bannerColor}
		t.Style().Color.Row = text.Colors{bannerColor}

		// Append the banner rows to the table
		t.AppendRow(table.Row{text.Bold.Sprint(sb.Config().Name)}, table.RowConfig{AutoMerge: true})
		t.AppendRow(table.Row{banner.Text})

		// Print the banner
		t.Render()
	}

}
