package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "Create CSV index for faster access",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(indexCmd).Standalone()

	indexCmd.Flags().StringP("delimiter", "d", "", "The field delimiter for reading CSV data.")
	indexCmd.Flags().StringP("output", "o", "", "Write index to <file> instead of <input>.idx.")
	rootCmd.AddCommand(indexCmd)
}
