package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/xsv_completer/cmd/common"
	"github.com/spf13/cobra"
)

var sortCmd = &cobra.Command{
	Use:   "sort",
	Short: "Sort CSV data",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sortCmd).Standalone()

	sortCmd.Flags().BoolP("numeric", "N", false, "Compare according to string numerical value")
	sortCmd.Flags().BoolP("reverse", "R", false, "Reverse order")
	sortCmd.Flags().StringP("select", "s", "", "Select a subset of columns to sort.")
	common.AddDelimiterFlag(sortCmd)
	common.AddNoHeadersFlag(sortCmd)
	common.AddOutputFlag(sortCmd)
	rootCmd.AddCommand(sortCmd)
}
