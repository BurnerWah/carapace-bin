package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upgradeMinimalCmd = &cobra.Command{
	Use:    "upgrade-minimal",
	Short:  "",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgradeMinimalCmd).Standalone()

	rootCmd.AddCommand(upgradeMinimalCmd)
}
