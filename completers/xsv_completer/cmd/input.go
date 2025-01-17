package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/xsv_completer/cmd/common"
	"github.com/spf13/cobra"
)

var inputCmd = &cobra.Command{
	Use:   "input",
	Short: "Read CSV data with special quoting rules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inputCmd).Standalone()

	inputCmd.Flags().String("escape", "", "The escape character to use. When not specified,")
	inputCmd.Flags().Bool("no-quoting", false, "Disable quoting completely.")
	inputCmd.Flags().String("quote", "", "The quote character to use. [default: \"]")
	common.AddDelimiterFlag(inputCmd)
	common.AddOutputFlag(inputCmd)
	rootCmd.AddCommand(inputCmd)
}
