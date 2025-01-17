package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "Concatenate by row or column",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(catCmd).Standalone()

	catCmd.Flags().StringP("delimiter", "d", "", "The field delimiter for reading CSV data")
	catCmd.Flags().BoolP("no-headers", "n", false, "Don't interpert the first row as a header")
	catCmd.Flags().StringP("output", "o", "", "Write output to a file instead of stdout")
	catCmd.Flags().BoolP("pad", "p", false, "Pad rows if data isn't long enough, or show all records for columns")
	rootCmd.AddCommand(catCmd)

	carapace.Gen(catCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
	})
	carapace.Gen(catCmd).PositionalCompletion(carapace.ActionValues("rows", "columns"))
	carapace.Gen(catCmd).PositionalAnyCompletion(carapace.ActionFiles(".csv"))
}
