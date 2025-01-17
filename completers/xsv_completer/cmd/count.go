package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var countCmd = &cobra.Command{
	Use:   "count",
	Short: "Count records",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(countCmd).Standalone()

	countCmd.Flags().StringP("delimiter", "d", "", "The field delimiter for reading CSV data.")
	countCmd.Flags().BoolP("no-headers", "n", false, "When set, the first row will not be included in")
	rootCmd.AddCommand(countCmd)
}
