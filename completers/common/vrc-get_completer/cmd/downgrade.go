package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var downgradeCmd = &cobra.Command{
	Use:   "downgrade",
	Short: "Downgrade the specified package specified version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(downgradeCmd).Standalone()

	downgradeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	downgradeCmd.Flags().Bool("no-update", false, "do not update local repository cache")
	downgradeCmd.Flags().Bool("offline", false, "do not connect to remote servers, use local caches only. implicitly --no-update")
	downgradeCmd.Flags().Bool("prerelease", false, "Include prerelease")
	downgradeCmd.Flags().StringP("project", "p", "", "Path to project dir. by default CWD or parents of CWD will be used")
	downgradeCmd.Flags().BoolP("version", "V", false, "Print version")
	downgradeCmd.Flags().BoolP("yes", "y", false, "skip confirm")
	rootCmd.AddCommand(downgradeCmd)
}
