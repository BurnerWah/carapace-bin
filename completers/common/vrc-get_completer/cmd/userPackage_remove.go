package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var userPackage_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove user package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(userPackage_removeCmd).Standalone()

	userPackage_removeCmd.Flags().BoolP("help", "h", false, "Print help")
	userPackage_removeCmd.Flags().BoolP("version", "V", false, "Print version")
	userPackageCmd.AddCommand(userPackage_removeCmd)
}
