package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var leavesCmd = &cobra.Command{
	Use:     "leaves",
	Short:   "List groups of installed packages not required by other installed packages",
	GroupID: "query",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(leavesCmd).Standalone()

	rootCmd.AddCommand(leavesCmd)
}
