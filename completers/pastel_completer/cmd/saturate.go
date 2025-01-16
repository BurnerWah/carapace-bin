package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var saturateCmd = &cobra.Command{
	Use:   "saturate",
	Short: "Increase color saturation by a specified amount",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(saturateCmd).Standalone()

	rootCmd.AddCommand(saturateCmd)
}
