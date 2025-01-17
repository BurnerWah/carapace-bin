package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(catCmd).Standalone()

	catCmd.Flags().StringP("delimiter", "d", "", "The field delimiter for reading CSV data.")
	catCmd.Flags().BoolP("help", "h", false, "Display this message")
	catCmd.Flags().BoolP("no-headers", "n", false, "When set, the first row will NOT be interpreted")
	catCmd.Flags().StringP("output", "o", "", "Write output to <file> instead of stdout.")
	catCmd.Flags().BoolP("pad", "p", false, "When concatenating columns, this flag will cause")
	rootCmd.AddCommand(catCmd)
}
