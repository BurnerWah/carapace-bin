package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var group_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display package list of a group",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(group_infoCmd).Standalone()

	group_infoCmd.Flags().Bool("available", false, "Show only available groups")
	group_infoCmd.Flags().StringSlice("contains-pkgs", nil, "Show only groups containing packages with specified names")
	group_infoCmd.Flags().Bool("hidden", false, "Show also hidden groups")
	group_infoCmd.Flags().Bool("installed", false, "Show only installed groups")
	groupCmd.AddCommand(group_infoCmd)
}
