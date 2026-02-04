package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var userPackage_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all user packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(userPackage_listCmd).Standalone()

	userPackage_listCmd.Flags().BoolP("help", "h", false, "Print help")
	userPackage_listCmd.Flags().BoolP("version", "V", false, "Print version")
	userPackageCmd.AddCommand(userPackage_listCmd)
}
