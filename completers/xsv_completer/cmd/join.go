package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/xsv_completer/cmd/common"
	"github.com/spf13/cobra"
)

var joinCmd = &cobra.Command{
	Use:   "join",
	Short: "Join CSV files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(joinCmd).Standalone()

	joinCmd.Flags().Bool("cross", false, "USE WITH CAUTION.")
	joinCmd.Flags().Bool("full", false, "Do a 'full outer' join. This returns all rows in")
	joinCmd.Flags().Bool("left", false, "Do a 'left outer' join. This returns all rows in")
	joinCmd.Flags().Bool("no-case", false, "When set, joins are done case insensitively.")
	joinCmd.Flags().Bool("nulls", false, "When set, joins will work on empty fields.")
	joinCmd.Flags().Bool("right", false, "Do a 'right outer' join. This returns all rows in")
	common.AddDelimiterFlag(joinCmd)
	common.AddNoHeadersFlag(joinCmd)
	common.AddOutputFlag(joinCmd)
	rootCmd.AddCommand(joinCmd)
}
