package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var flattenCmd = &cobra.Command{
	Use:   "flatten",
	Short: "Show one field per line",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flattenCmd).Standalone()

	flattenCmd.Flags().StringP("condense", "c", "", "Limits the length of each field to the value")
	flattenCmd.Flags().StringP("delimiter", "d", "", "The field delimiter for reading CSV data.")
	flattenCmd.Flags().BoolP("no-headers", "n", false, "When set, the first row will not be interpreted")
	flattenCmd.Flags().StringP("separator", "s", "", "A string of characters to write after each record.")
	rootCmd.AddCommand(flattenCmd)
}
