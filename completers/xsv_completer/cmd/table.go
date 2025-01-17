package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/xsv_completer/cmd/common"
	"github.com/spf13/cobra"
)

var tableCmd = &cobra.Command{
	Use:   "table",
	Short: "Align CSV data into columns",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tableCmd).Standalone()

	tableCmd.Flags().StringP("condense", "c", "", "Limits the length of each field to the value")
	tableCmd.Flags().StringP("pad", "p", "", "The minimum number of spaces between each column.")
	tableCmd.Flags().StringP("width", "w", "", "The minimum width of each column.")
	common.AddDelimiterFlag(tableCmd)
	common.AddOutputFlag(tableCmd)
	rootCmd.AddCommand(tableCmd)
}
