package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_reinstallCmd = &cobra.Command{
	Use:   "reinstall",
	Short: "Reinstall specified packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_reinstallCmd).Standalone()

	helpCmd.AddCommand(help_reinstallCmd)
}
