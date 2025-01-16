package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rotateCmd = &cobra.Command{
	Use:   "rotate",
	Short: "Rotate the hue channel by the specified angle",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rotateCmd).Standalone()

	rootCmd.AddCommand(rotateCmd)
}
