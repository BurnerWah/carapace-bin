package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inputCmd = &cobra.Command{
	Use:   "input",
	Short: "Read CSV data with special quoting rules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inputCmd).Standalone()

	inputCmd.Flags().StringP("delimiter", "d", "", "The field delimiter for reading CSV data.")
	inputCmd.Flags().String("escape", "", "The escape character to use. When not specified,")
	inputCmd.Flags().Bool("no-quoting", false, "Disable quoting completely.")
	inputCmd.Flags().StringP("output", "o", "", "Write output to <file> instead of stdout.")
	inputCmd.Flags().String("quote", "", "The quote character to use. [default: \"]")
	rootCmd.AddCommand(inputCmd)
}
