package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install",
	Short:   "Adds package to unity project",
	Aliases: []string{"i"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	installCmd.Flags().BoolP("name", "n", false, "Install package by display name instead of name")
	installCmd.Flags().Bool("no-update", false, "do not update local repository cache")
	installCmd.Flags().Bool("offline", false, "do not connect to remote servers, use local caches only. implicitly --no-update")
	installCmd.Flags().Bool("prerelease", false, "Include prerelease")
	installCmd.Flags().StringP("project", "p", "", "Path to project dir. by default CWD or parents of CWD will be used")
	installCmd.Flags().BoolP("version", "V", false, "Print version")
	installCmd.Flags().BoolP("yes", "y", false, "skip confirm")
	rootCmd.AddCommand(installCmd)
}
