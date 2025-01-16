package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var colorCmd = &cobra.Command{
	Use:   "color",
	Short: "Display information about the given color",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(colorCmd).Standalone()

	rootCmd.AddCommand(colorCmd)
}
