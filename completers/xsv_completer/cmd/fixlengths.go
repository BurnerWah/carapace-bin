package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/xsv_completer/cmd/common"
	"github.com/spf13/cobra"
)

var fixlengthsCmd = &cobra.Command{
	Use:   "fixlengths",
	Short: "Makes all records have same length",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fixlengthsCmd).Standalone()

	fixlengthsCmd.Flags().StringP("length", "l", "", "Forcefully set the length of each record. If a")
	common.AddDelimiterFlag(fixlengthsCmd)
	common.AddOutputFlag(fixlengthsCmd)
	rootCmd.AddCommand(fixlengthsCmd)
}
