package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrate_eslintCmd = &cobra.Command{
	Use:   "eslint",
	Short: "Migrate an ESLint configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrate_eslintCmd).Standalone()

	migrate_eslintCmd.Flags().Bool("include-inspired", false, "Includes rules inspired from an eslint rule in the migration")
	migrate_eslintCmd.Flags().Bool("include-nursery", false, "Includes nursery rules in the migration")
	migrateCmd.AddCommand(migrate_eslintCmd)
}
