package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:     "info",
	Short:   "Describes the given package",
	GroupID: "query",
	Aliases: []string{"if"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()

	infoCmd.Flags().Bool("autoremove", false, "List packages which will be removed by the 'dnf autoremove' command")
	infoCmd.Flags().Bool("available", false, "List available packages")
	infoCmd.Flags().Bool("extras", false, "List installed packages that are not available in any repository")
	infoCmd.Flags().Bool("installed", false, "List installed packages")
	infoCmd.Flags().Bool("obsoletes", false, "List installed packages that are obsoleted by packages in any repository")
	infoCmd.Flags().Bool("recent", false, "List packages recently added into the repositories")
	infoCmd.Flags().Bool("showduplicates", false, "Show all versions of the packages, not only the latest ones")
	infoCmd.Flags().Bool("updates", false, "Alias for '--upgrades'")
	infoCmd.Flags().Bool("upgrades", false, "List upgrades available for the installed packages")
	infoCmd.Flag("updates").Hidden = true
	rootCmd.AddCommand(infoCmd)
}
