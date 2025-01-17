package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/xsv_completer/cmd/common"
	"github.com/spf13/cobra"
)

var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "Select columns from CSV",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(selectCmd).Standalone()

	common.AddDelimiterFlag(selectCmd)
	common.AddNoHeadersFlag(selectCmd)
	common.AddOutputFlag(selectCmd)
	rootCmd.AddCommand(selectCmd)
}
