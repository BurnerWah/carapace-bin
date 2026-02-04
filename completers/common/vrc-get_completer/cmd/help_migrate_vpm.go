package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_migrate_vpmCmd = &cobra.Command{
	Use:   "vpm",
	Short: "Migrate your legacy (unitypackage) VRCSDK project to VPM project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_migrate_vpmCmd).Standalone()

	help_migrateCmd.AddCommand(help_migrate_vpmCmd)
}
