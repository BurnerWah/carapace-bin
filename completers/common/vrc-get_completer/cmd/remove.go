package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove package from Unity project",
	Aliases: []string{"rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeCmd).Standalone()

	removeCmd.Flags().BoolP("help", "h", false, "Print help")
	removeCmd.Flags().Bool("no-update", false, "do not update local repository cache")
	removeCmd.Flags().Bool("offline", false, "do not connect to remote servers, use local caches only. implicitly --no-update")
	removeCmd.Flags().StringP("project", "p", "", "Path to project dir. by default CWD or parents of CWD will be used")
	removeCmd.Flags().BoolP("version", "V", false, "Print version")
	removeCmd.Flags().BoolP("yes", "y", false, "skip confirm")
	rootCmd.AddCommand(removeCmd)
}
