package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configManager_setoptCmd = &cobra.Command{
	Use:   "setopt",
	Short: "Set configuration and repositories options",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configManager_setoptCmd).Standalone()

	configManager_setoptCmd.Flags().Bool("create-missing-dir", false, "Allow creation of missing directories")
	configManagerCmd.AddCommand(configManager_setoptCmd)
}
