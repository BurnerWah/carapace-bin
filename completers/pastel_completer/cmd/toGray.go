package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var toGrayCmd = &cobra.Command{
	Use:   "to-gray",
	Short: "Completely desaturate a color (preserving luminance)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(toGrayCmd).Standalone()

	rootCmd.AddCommand(toGrayCmd)
}
