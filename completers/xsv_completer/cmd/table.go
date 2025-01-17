package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tableCmd = &cobra.Command{
	Use:   "table",
	Short: "Align CSV data into columns",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tableCmd).Standalone()

	tableCmd.Flags().StringP("condense", "c", "", "Limits the length of each field to the value")
	tableCmd.Flags().StringP("delimiter", "d", "", "The field delimiter for reading CSV data.")
	tableCmd.Flags().StringP("output", "o", "", "Write output to <file> instead of stdout.")
	tableCmd.Flags().StringP("pad", "p", "", "The minimum number of spaces between each column.")
	tableCmd.Flags().StringP("width", "w", "", "The minimum width of each column.")
	rootCmd.AddCommand(tableCmd)
}
