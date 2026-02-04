package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade specified package or all packages to latest or specified version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgradeCmd).Standalone()

	upgradeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	upgradeCmd.Flags().Bool("no-update", false, "do not update local repository cache")
	upgradeCmd.Flags().Bool("offline", false, "do not connect to remote servers, use local caches only. implicitly --no-update")
	upgradeCmd.Flags().Bool("prerelease", false, "Include prerelease")
	upgradeCmd.Flags().StringP("project", "p", "", "Path to project dir. by default CWD or parents of CWD will be used")
	upgradeCmd.Flags().BoolP("version", "V", false, "Print version")
	upgradeCmd.Flags().BoolP("yes", "y", false, "skip confirm")
	rootCmd.AddCommand(upgradeCmd)
}
