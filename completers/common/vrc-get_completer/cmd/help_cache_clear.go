package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_cache_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Cleanup package cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_cache_clearCmd).Standalone()

	help_cacheCmd.AddCommand(help_cache_clearCmd)
}
