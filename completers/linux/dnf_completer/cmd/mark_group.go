package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mark_groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Mark package as installed by a group",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mark_groupCmd).Standalone()

	markCmd.AddCommand(mark_groupCmd)
}
