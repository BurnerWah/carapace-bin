package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outdatedCmd = &cobra.Command{
	Use:   "outdated",
	Short: "Show list of outdated packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outdatedCmd).Standalone()

	outdatedCmd.Flags().BoolP("help", "h", false, "Print help")
	outdatedCmd.Flags().String("json-format", "", "With this option, output is printed in json format")
	outdatedCmd.Flags().Bool("no-update", false, "do not update local repository cache")
	outdatedCmd.Flags().Bool("offline", false, "do not connect to remote servers, use local caches only. implicitly --no-update")
	outdatedCmd.Flags().Bool("prerelease", false, "Include prerelease")
	outdatedCmd.Flags().StringP("project", "p", "", "Path to project dir. by default CWD or parents of CWD will be used")
	outdatedCmd.Flags().BoolP("version", "V", false, "Print version")
	rootCmd.AddCommand(outdatedCmd)
}
