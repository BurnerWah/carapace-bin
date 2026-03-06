package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configManager_unsetoptCmd = &cobra.Command{
	Use:   "unsetopt",
	Short: "Unset/remove configuration and repositories options",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configManager_unsetoptCmd).Standalone()

	configManagerCmd.AddCommand(configManager_unsetoptCmd)
}
