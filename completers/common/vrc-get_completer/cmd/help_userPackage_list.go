package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_userPackage_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all user packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_userPackage_listCmd).Standalone()

	help_userPackageCmd.AddCommand(help_userPackage_listCmd)
}
