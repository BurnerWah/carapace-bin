package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var desaturateCmd = &cobra.Command{
	Use:   "desaturate",
	Short: "Decrease color saturation by a specified amount",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(desaturateCmd).Standalone()

	rootCmd.AddCommand(desaturateCmd)
}
