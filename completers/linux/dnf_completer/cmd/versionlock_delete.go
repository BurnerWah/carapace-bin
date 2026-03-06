package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionlock_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove any matching versionlock config entries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionlock_deleteCmd).Standalone()

	versionlockCmd.AddCommand(versionlock_deleteCmd)
}
