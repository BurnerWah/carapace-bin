package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/xsv_completer/cmd/common"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search CSV data with regexes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	searchCmd.Flags().BoolP("ignore-case", "i", false, "Case insensitive search. This is equivalent to")
	searchCmd.Flags().BoolP("invert-match", "v", false, "Select only rows that did not match")
	searchCmd.Flags().StringP("select", "s", "", "Select the columns to search. See 'xsv select -h'")
	common.AddDelimiterFlag(searchCmd)
	common.AddNoHeadersFlag(searchCmd)
	common.AddOutputFlag(searchCmd)
	rootCmd.AddCommand(searchCmd)
}
