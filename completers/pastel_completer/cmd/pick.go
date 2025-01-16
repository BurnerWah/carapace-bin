package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pickCmd = &cobra.Command{
	Use:   "pick",
	Short: "Interactively pick a color from the screen (pipette)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pickCmd).Standalone()

	rootCmd.AddCommand(pickCmd)
}
