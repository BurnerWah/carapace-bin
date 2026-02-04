package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrate_vpmCmd = &cobra.Command{
	Use:   "vpm",
	Short: "Migrate your legacy (unitypackage) VRCSDK project to VPM project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrate_vpmCmd).Standalone()

	migrate_vpmCmd.Flags().BoolP("help", "h", false, "Print help")
	migrate_vpmCmd.Flags().Bool("no-update", false, "do not update local repository cache")
	migrate_vpmCmd.Flags().Bool("offline", false, "do not connect to remote servers, use local caches only. implicitly --no-update")
	migrate_vpmCmd.Flags().StringP("project", "p", "", "Path to project dir. by default CWD or parents of CWD will be used")
	migrateCmd.AddCommand(migrate_vpmCmd)
}
