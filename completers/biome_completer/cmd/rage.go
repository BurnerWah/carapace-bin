package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/biome_completer/cmd/common"
	"github.com/spf13/cobra"
)

var rageCmd = &cobra.Command{
	Use:   "rage",
	Short: "Prints information for debugging",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rageCmd).Standalone()

	rageCmd.Flags().Bool("daemon-logs", false, "Prints the Biome daemon server logs")
	rageCmd.Flags().Bool("formatter", false, "Prints the formatter options applied")
	rageCmd.Flags().Bool("linter", false, "Prints the linter options applied")
	common.AddGlobalFlags(rageCmd)
	rootCmd.AddCommand(rageCmd)
}
