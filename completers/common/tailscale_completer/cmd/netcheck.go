package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var netcheckCmd = &cobra.Command{
	Use:   "netcheck",
	Short: "Print an analysis of local network conditions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(netcheckCmd).Standalone()

	netcheckCmd.Flags().Duration("every", 0, "if non-zero, do an incremental report with the given frequency")
	netcheckCmd.Flags().String("format", "", "output format")
	netcheckCmd.Flags().Bool("verbose", false, "verbose logs")

	rootCmd.AddCommand(netcheckCmd)

	carapace.Gen(netcheckCmd).FlagCompletion(carapace.ActionMap{
		// This is goofy but how the command works
		"format": carapace.ActionValuesDescribed(
			"", "human-readable",
			"json", "",
			"json-line", "",
		),
	})
}
