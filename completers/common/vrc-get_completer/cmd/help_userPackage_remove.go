package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_userPackage_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove user package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_userPackage_removeCmd).Standalone()

	help_userPackageCmd.AddCommand(help_userPackage_removeCmd)
}
