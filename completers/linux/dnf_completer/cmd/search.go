package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:     "search",
	Short:   "Search for software matching all specified strings",
	GroupID: "query",
	Aliases: []string{"se"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	searchCmd.Flags().Bool("all", false, "Search also package description and URL")
	searchCmd.Flags().Bool("showduplicates", false, "Show all versions of the packages, not only the latest ones")
	rootCmd.AddCommand(searchCmd)
}
