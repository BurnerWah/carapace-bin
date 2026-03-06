package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "Lists packages depending on the packages' relation to the system",
	GroupID: "query",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().Bool("autoremove", false, "List packages which will be removed by the 'dnf autoremove' command")
	listCmd.Flags().Bool("available", false, "List available packages")
	listCmd.Flags().Bool("extras", false, "List installed packages that are not available in any repository")
	listCmd.Flags().Bool("installed", false, "List installed packages")
	listCmd.Flags().Bool("obsoletes", false, "List installed packages that are obsoleted by packages in any repository")
	listCmd.Flags().Bool("recent", false, "List packages recently added into the repositories")
	listCmd.Flags().Bool("showduplicates", false, "Show all versions of the packages, not only the latest ones")
	listCmd.Flags().Bool("updates", false, "Alias for '--upgrades'")
	listCmd.Flags().Bool("upgrades", false, "List upgrades available for the installed packages")
	listCmd.Flag("updates").Hidden = true
	rootCmd.AddCommand(listCmd)
}
