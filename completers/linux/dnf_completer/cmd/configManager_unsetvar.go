package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configManager_unsetvarCmd = &cobra.Command{
	Use:   "unsetvar",
	Short: "Unset/remove variables",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configManager_unsetvarCmd).Standalone()

	configManagerCmd.AddCommand(configManager_unsetvarCmd)
}
