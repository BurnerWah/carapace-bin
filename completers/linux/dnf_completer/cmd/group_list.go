package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var group_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List groups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(group_listCmd).Standalone()

	group_listCmd.Flags().Bool("available", false, "Show only available groups")
	group_listCmd.Flags().StringSlice("contains-pkgs", nil, "Show only groups containing packages with specified names")
	group_listCmd.Flags().Bool("hidden", false, "Show also hidden groups")
	group_listCmd.Flags().Bool("installed", false, "Show only installed groups")
	groupCmd.AddCommand(group_listCmd)
}
