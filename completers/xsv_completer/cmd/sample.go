package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sampleCmd = &cobra.Command{
	Use:   "sample",
	Short: "Randomly sample CSV data",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sampleCmd).Standalone()

	sampleCmd.Flags().StringP("delimiter", "d", "", "The field delimiter for reading CSV data.")
	sampleCmd.Flags().BoolP("help", "h", false, "Display this message")
	sampleCmd.Flags().BoolP("no-headers", "n", false, "When set, the first row will be consider as part of")
	sampleCmd.Flags().StringP("output", "o", "", "Write output to <file> instead of stdout.")
	rootCmd.AddCommand(sampleCmd)
}
