package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_help_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Cleanup package cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_help_clearCmd).Standalone()

	cache_helpCmd.AddCommand(cache_help_clearCmd)
}
