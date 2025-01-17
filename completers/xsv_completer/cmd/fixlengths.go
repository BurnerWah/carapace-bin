package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fixlengthsCmd = &cobra.Command{
	Use:   "fixlengths",
	Short: "Makes all records have same length",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fixlengthsCmd).Standalone()

	fixlengthsCmd.Flags().StringP("delimiter", "d", "", "The field delimiter for reading CSV data.")
	fixlengthsCmd.Flags().StringP("length", "l", "", "Forcefully set the length of each record. If a")
	fixlengthsCmd.Flags().StringP("output", "o", "", "Write output to <file> instead of stdout.")
	rootCmd.AddCommand(fixlengthsCmd)
}
