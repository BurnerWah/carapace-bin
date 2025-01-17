package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sortCmd = &cobra.Command{
	Use:   "sort",
	Short: "Sort CSV data",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sortCmd).Standalone()

	sortCmd.Flags().StringP("delimiter", "d", "", "The field delimiter for reading CSV data.")
	sortCmd.Flags().BoolP("help", "h", false, "Display this message")
	sortCmd.Flags().BoolP("no-headers", "n", false, "When set, the first row will not be interpreted")
	sortCmd.Flags().BoolP("numeric", "N", false, "Compare according to string numerical value")
	sortCmd.Flags().StringP("output", "o", "", "Write output to <file> instead of stdout.")
	sortCmd.Flags().BoolP("reverse", "R", false, "Reverse order")
	sortCmd.Flags().StringP("select", "s", "", "Select a subset of columns to sort.")
	rootCmd.AddCommand(sortCmd)
}
