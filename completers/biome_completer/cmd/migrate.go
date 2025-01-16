package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/biome_completer/cmd/common"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Updates the configuration when there are breaking changes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrateCmd).Standalone()

	migrateCmd.Flags().Bool("fix", false, "Alias of `--write`, writes the new configuration file to disk")
	migrateCmd.Flags().Bool("write", false, "Writes the new configuration file to disk")
	common.AddGlobalFlags(migrateCmd)
	rootCmd.AddCommand(migrateCmd)
}
