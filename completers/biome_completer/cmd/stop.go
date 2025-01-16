package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops the Biome daemon server process",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stopCmd).Standalone()

	rootCmd.AddCommand(stopCmd)
}
