package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Commands about cache control",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_cacheCmd).Standalone()

	helpCmd.AddCommand(help_cacheCmd)
}
