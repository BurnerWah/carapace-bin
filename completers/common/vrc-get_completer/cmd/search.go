package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search package by the query",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	searchCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	searchCmd.Flags().Bool("no-update", false, "do not update local repository cache")
	searchCmd.Flags().Bool("offline", false, "do not connect to remote servers, use local caches only. implicitly --no-update")
	searchCmd.Flags().BoolP("version", "V", false, "Print version")
	rootCmd.AddCommand(searchCmd)
}
