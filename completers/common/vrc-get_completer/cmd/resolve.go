package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resolveCmd = &cobra.Command{
	Use:   "resolve",
	Short: "(re)installs all locked packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resolveCmd).Standalone()

	resolveCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	resolveCmd.Flags().Bool("no-update", false, "do not update local repository cache")
	resolveCmd.Flags().Bool("offline", false, "do not connect to remote servers, use local caches only. implicitly --no-update")
	resolveCmd.Flags().StringP("project", "p", "", "Path to project dir. by default CWD or parents of CWD will be used")
	resolveCmd.Flags().BoolP("version", "V", false, "Print version")
	rootCmd.AddCommand(resolveCmd)
}
