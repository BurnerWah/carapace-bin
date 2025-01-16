package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var grayCmd = &cobra.Command{
	Use:   "gray",
	Short: "Create a gray tone from a given lightness",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(grayCmd).Standalone()

	rootCmd.AddCommand(grayCmd)
}
