package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/xsv_completer/cmd/common"
	"github.com/spf13/cobra"
)

var sampleCmd = &cobra.Command{
	Use:   "sample",
	Short: "Randomly sample CSV data",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sampleCmd).Standalone()

	common.AddDelimiterFlag(sampleCmd)
	common.AddNoHeadersFlag(sampleCmd)
	common.AddOutputFlag(sampleCmd)
	rootCmd.AddCommand(sampleCmd)
}
