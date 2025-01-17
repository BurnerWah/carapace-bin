package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/xsv_completer/cmd/common"
	"github.com/spf13/cobra"
)

var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "Concatenate by row or column",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(catCmd).Standalone()

	catCmd.Flags().BoolP("pad", "p", false, "Pad rows if data isn't long enough, or show all records for columns")
	common.AddCommonFlags(catCmd)
	common.AddOutputFlag(catCmd)
	rootCmd.AddCommand(catCmd)

	carapace.Gen(catCmd).PositionalCompletion(carapace.ActionValues("rows", "columns"))
	carapace.Gen(catCmd).PositionalAnyCompletion(carapace.ActionFiles(".csv"))
}
