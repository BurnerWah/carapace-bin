package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search CSV data with regexes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	searchCmd.Flags().StringP("delimiter", "d", "", "The field delimiter for reading CSV data.")
	searchCmd.Flags().BoolP("ignore-case", "i", false, "Case insensitive search. This is equivalent to")
	searchCmd.Flags().BoolP("invert-match", "v", false, "Select only rows that did not match")
	searchCmd.Flags().BoolP("no-headers", "n", false, "When set, the first row will not be interpreted")
	searchCmd.Flags().StringP("output", "o", "", "Write output to <file> instead of stdout.")
	searchCmd.Flags().StringP("select", "s", "", "Select the columns to search. See 'xsv select -h'")
	rootCmd.AddCommand(searchCmd)
}
