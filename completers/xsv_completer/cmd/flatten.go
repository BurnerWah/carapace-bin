package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/xsv_completer/cmd/common"
	"github.com/spf13/cobra"
)

var flattenCmd = &cobra.Command{
	Use:   "flatten",
	Short: "Show one field per line",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flattenCmd).Standalone()

	flattenCmd.Flags().StringP("condense", "c", "", "Limits the length of each field to the value")
	flattenCmd.Flags().StringP("separator", "s", "", "A string of characters to write after each record.")
	common.AddDelimiterFlag(flattenCmd)
	common.AddNoHeadersFlag(flattenCmd)
	rootCmd.AddCommand(flattenCmd)
}
