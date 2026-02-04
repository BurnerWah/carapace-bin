package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Cleanup package cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_clearCmd).Standalone()

	cache_clearCmd.Flags().BoolP("help", "h", false, "Print help")
	cache_clearCmd.Flags().Bool("no-update", false, "do not update local repository cache")
	cache_clearCmd.Flags().Bool("offline", false, "do not connect to remote servers, use local caches only. implicitly --no-update")
	cache_clearCmd.Flags().BoolP("version", "V", false, "Print version")
	cacheCmd.AddCommand(cache_clearCmd)
}
