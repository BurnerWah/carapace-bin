package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var userPackageCmd = &cobra.Command{
	Use:   "user-package",
	Short: "Commands around user packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(userPackageCmd).Standalone()

	userPackageCmd.Flags().BoolP("help", "h", false, "Print help")
	userPackageCmd.Flags().BoolP("version", "V", false, "Print version")
	rootCmd.AddCommand(userPackageCmd)
}
