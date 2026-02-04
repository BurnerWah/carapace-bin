package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate Unity Project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrateCmd).Standalone()

	migrateCmd.Flags().BoolP("help", "h", false, "Print help")
	migrateCmd.Flags().BoolP("version", "V", false, "Print version")
	rootCmd.AddCommand(migrateCmd)
}
