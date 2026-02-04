package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrate_unity2022Cmd = &cobra.Command{
	Use:   "unity2022",
	Short: "Migrate your project to Unity 2022",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrate_unity2022Cmd).Standalone()

	migrate_unity2022Cmd.Flags().BoolP("help", "h", false, "Print help")
	migrate_unity2022Cmd.Flags().Bool("no-update", false, "do not update local repository cache")
	migrate_unity2022Cmd.Flags().Bool("offline", false, "do not connect to remote servers, use local caches only. implicitly --no-update")
	migrate_unity2022Cmd.Flags().StringP("project", "p", "", "Path to project dir. by default CWD or parents of CWD will be used")
	migrate_unity2022Cmd.Flags().String("unity", "", "Path to unity 2022 executable")
	migrate_unity2022Cmd.MarkFlagRequired("unity")
	migrateCmd.AddCommand(migrate_unity2022Cmd)
}
