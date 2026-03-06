package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionlockCmd = &cobra.Command{
	Use:     "versionlock",
	Short:   "Manage versionlock config",
	GroupID: "subcommands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionlockCmd).Standalone()

	rootCmd.AddCommand(versionlockCmd)
}
