package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var headersCmd = &cobra.Command{
	Use:   "headers",
	Short: "Show header names",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(headersCmd).Standalone()

	headersCmd.Flags().StringP("delimiter", "d", "", "The field delimiter for reading CSV data.")
	headersCmd.Flags().Bool("intersect", false, "Shows the intersection of all headers in all of")
	headersCmd.Flags().BoolP("just-names", "j", false, "Only show the header names (hide column index).")
	rootCmd.AddCommand(headersCmd)
}
