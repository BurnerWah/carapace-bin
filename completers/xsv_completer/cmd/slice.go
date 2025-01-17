package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sliceCmd = &cobra.Command{
	Use:   "slice",
	Short: "Slice records from CSV",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sliceCmd).Standalone()

	sliceCmd.Flags().StringP("delimiter", "d", "", "The field delimiter for reading CSV data.")
	sliceCmd.Flags().StringP("end", "e", "", "The index of the record to slice to.")
	sliceCmd.Flags().BoolP("help", "h", false, "Display this message")
	sliceCmd.Flags().StringP("index", "i", "", "Slice a single record (shortcut for -s N -l 1).")
	sliceCmd.Flags().StringP("len", "l", "", "The length of the slice (can be used instead")
	sliceCmd.Flags().BoolP("no-headers", "n", false, "When set, the first row will not be interpreted")
	sliceCmd.Flags().StringP("output", "o", "", "Write output to <file> instead of stdout.")
	sliceCmd.Flags().StringP("start", "s", "", "The index of the record to slice from.")
	rootCmd.AddCommand(sliceCmd)
}
