package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/biome_completer/cmd/common"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Searches for Grit patterns across a project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	common.AddGlobalFlags(searchCmd)
	common.AddVcsFlags(searchCmd)
	common.AddFilesystemFlags(searchCmd)
	common.AddStdinFilePathFlag(searchCmd)
	rootCmd.AddCommand(searchCmd)
}
