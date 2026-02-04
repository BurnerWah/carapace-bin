package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var info_packageCmd = &cobra.Command{
	Use:   "package",
	Short: "Show project information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(info_packageCmd).Standalone()

	info_packageCmd.Flags().BoolP("help", "h", false, "Print help")
	info_packageCmd.Flags().String("json-format", "", "Output json format")
	info_packageCmd.Flags().Bool("no-update", false, "do not update local repository cache")
	info_packageCmd.Flags().Bool("offline", false, "do not connect to remote servers, use local caches only. implicitly --no-update")
	info_packageCmd.Flags().BoolP("version", "V", false, "Print version")
	infoCmd.AddCommand(info_packageCmd)
}
