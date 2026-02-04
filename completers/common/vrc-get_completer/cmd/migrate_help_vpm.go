package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrate_help_vpmCmd = &cobra.Command{
	Use:   "vpm",
	Short: "Migrate your legacy (unitypackage) VRCSDK project to VPM project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrate_help_vpmCmd).Standalone()

	migrate_helpCmd.AddCommand(migrate_help_vpmCmd)
}
