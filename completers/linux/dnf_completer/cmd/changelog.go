package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var changelogCmd = &cobra.Command{
	Use:     "changelog",
	Short:   "Show package changelogs",
	GroupID: "commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(changelogCmd).Standalone()

	changelogCmd.Flags().String("count", "", "Limit the number of changelog entries shown per package")
	changelogCmd.Flags().String("since", "", "Show changelog entries since date in the YYYY-MM-DD format")
	changelogCmd.Flags().Bool("upgrades", false, "Show new changelog entries for packages that provide an upgrade for an already installed package")
	rootCmd.AddCommand(changelogCmd)
}
