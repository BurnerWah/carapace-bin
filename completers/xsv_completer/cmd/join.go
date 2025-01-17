package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var joinCmd = &cobra.Command{
	Use:   "join",
	Short: "Join CSV files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(joinCmd).Standalone()

	joinCmd.Flags().Bool("cross", false, "USE WITH CAUTION.")
	joinCmd.Flags().StringP("delimiter", "d", "", "The field delimiter for reading CSV data.")
	joinCmd.Flags().Bool("full", false, "Do a 'full outer' join. This returns all rows in")
	joinCmd.Flags().BoolP("help", "h", false, "Display this message")
	joinCmd.Flags().Bool("left", false, "Do a 'left outer' join. This returns all rows in")
	joinCmd.Flags().Bool("no-case", false, "When set, joins are done case insensitively.")
	joinCmd.Flags().BoolP("no-headers", "n", false, "When set, the first row will not be interpreted")
	joinCmd.Flags().Bool("nulls", false, "When set, joins will work on empty fields.")
	joinCmd.Flags().StringP("output", "o", "", "Write output to <file> instead of stdout.")
	joinCmd.Flags().Bool("right", false, "Do a 'right outer' join. This returns all rows in")
	rootCmd.AddCommand(joinCmd)
}
