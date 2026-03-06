package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var systemUpgrade_downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download everything needed to upgrade to a new release",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(systemUpgrade_downloadCmd).Standalone()

	systemUpgrade_downloadCmd.Flags().Bool("no-downgrade", false, "Do not install packages from the new release if they are older than what is currently installed")
	systemUpgradeCmd.AddCommand(systemUpgrade_downloadCmd)
}
