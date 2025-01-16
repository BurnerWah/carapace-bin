package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrate_prettierCmd = &cobra.Command{
	Use:   "prettier",
	Short: "Migrate a Prettier configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrate_prettierCmd).Standalone()

	migrateCmd.AddCommand(migrate_prettierCmd)
}
