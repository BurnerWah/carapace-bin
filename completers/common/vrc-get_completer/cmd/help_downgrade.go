package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_downgradeCmd = &cobra.Command{
	Use:   "downgrade",
	Short: "Downgrade the specified package specified version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_downgradeCmd).Standalone()

	helpCmd.AddCommand(help_downgradeCmd)
}
