package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var userPackage_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add user package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(userPackage_addCmd).Standalone()

	userPackage_addCmd.Flags().BoolP("help", "h", false, "Print help")
	userPackage_addCmd.Flags().BoolP("version", "V", false, "Print version")
	userPackageCmd.AddCommand(userPackage_addCmd)
}
