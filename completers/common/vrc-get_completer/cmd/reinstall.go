package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var reinstallCmd = &cobra.Command{
	Use:   "reinstall",
	Short: "Reinstall specified packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reinstallCmd).Standalone()

	reinstallCmd.Flags().BoolP("help", "h", false, "Print help")
	reinstallCmd.Flags().Bool("no-update", false, "do not update local repository cache")
	reinstallCmd.Flags().Bool("offline", false, "do not connect to remote servers, use local caches only. implicitly --no-update")
	reinstallCmd.Flags().StringP("project", "p", "", "Path to project dir. by default CWD or parents of CWD will be used")
	reinstallCmd.Flags().BoolP("version", "V", false, "Print version")
	reinstallCmd.Flags().BoolP("yes", "y", false, "skip confirm")
	rootCmd.AddCommand(reinstallCmd)
}
