package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/xsv_completer/cmd/common"
	"github.com/spf13/cobra"
)

var countCmd = &cobra.Command{
	Use:   "count",
	Short: "Count records",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(countCmd).Standalone()

	common.AddDelimiterFlag(countCmd)
	common.AddNoHeadersFlag(countCmd)
	rootCmd.AddCommand(countCmd)

	carapace.Gen(countCmd).PositionalAnyCompletion(carapace.ActionFiles(".csv"))
}
