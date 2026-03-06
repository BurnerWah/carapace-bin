package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configManager_addrepoCmd = &cobra.Command{
	Use:   "addrepo",
	Short: "Add repositories from the specified configuration file or define a new repository using user options",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configManager_addrepoCmd).Standalone()

	configManager_addrepoCmd.Flags().Bool("add-or-replace", false, "Allow adding or replacing a repository in the existing configuration file")
	configManager_addrepoCmd.Flags().Bool("create-missing-dir", false, "Allow creation of missing directories")
	configManager_addrepoCmd.Flags().String("from-repofile", "", "Download repository configuration file, test it and put it in reposdir")
	configManager_addrepoCmd.Flags().String("id", "", "Set id for newly created repository")
	configManager_addrepoCmd.Flags().Bool("overwrite", false, "Allow overwriting of existing repository configuration file")
	configManager_addrepoCmd.Flags().String("save-filename", "", "Set the name of the configuration file of the added repository")
	configManager_addrepoCmd.Flags().String("set", "", "Set option in newly created repository")
	configManagerCmd.AddCommand(configManager_addrepoCmd)
}
