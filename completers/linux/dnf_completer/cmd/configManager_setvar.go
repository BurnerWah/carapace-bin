package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configManager_setvarCmd = &cobra.Command{
	Use:   "setvar",
	Short: "Set variables",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configManager_setvarCmd).Standalone()

	configManager_setvarCmd.Flags().Bool("create-missing-dir", false, "Allow creation of missing directories")
	configManagerCmd.AddCommand(configManager_setvarCmd)
}
