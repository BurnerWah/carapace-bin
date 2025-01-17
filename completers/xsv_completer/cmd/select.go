package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(selectCmd).Standalone()

	selectCmd.Flags().StringP("delimiter", "d", "", "The field delimiter for reading CSV data.")
	selectCmd.Flags().BoolP("help", "h", false, "Display this message")
	selectCmd.Flags().BoolP("no-headers", "n", false, "When set, the first row will not be interpreted")
	selectCmd.Flags().StringP("output", "o", "", "Write output to <file> instead of stdout.")
	rootCmd.AddCommand(selectCmd)
}
