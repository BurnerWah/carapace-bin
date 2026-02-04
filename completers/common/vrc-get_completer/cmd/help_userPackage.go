package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_userPackageCmd = &cobra.Command{
	Use:   "user-package",
	Short: "Commands around user packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_userPackageCmd).Standalone()

	helpCmd.AddCommand(help_userPackageCmd)
}
