package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrate_help_unity2022Cmd = &cobra.Command{
	Use:   "unity2022",
	Short: "Migrate your project to Unity 2022",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrate_help_unity2022Cmd).Standalone()

	migrate_helpCmd.AddCommand(migrate_help_unity2022Cmd)
}
