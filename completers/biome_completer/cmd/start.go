package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/biome_completer/cmd/common"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the Biome daemon server process",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(startCmd).Standalone()

	common.AddServerFlags(startCmd)
	rootCmd.AddCommand(startCmd)
}
