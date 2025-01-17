package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/xsv_completer/cmd/common"
	"github.com/spf13/cobra"
)

var sliceCmd = &cobra.Command{
	Use:   "slice",
	Short: "Slice records from CSV",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sliceCmd).Standalone()

	sliceCmd.Flags().StringP("end", "e", "", "The index of the record to slice to.")
	sliceCmd.Flags().StringP("index", "i", "", "Slice a single record (shortcut for -s N -l 1).")
	sliceCmd.Flags().StringP("len", "l", "", "The length of the slice (can be used instead")
	sliceCmd.Flags().StringP("start", "s", "", "The index of the record to slice from.")
	common.AddDelimiterFlag(sliceCmd)
	common.AddNoHeadersFlag(sliceCmd)
	common.AddOutputFlag(sliceCmd)
	rootCmd.AddCommand(sliceCmd)
}
