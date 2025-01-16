package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a color property to a specific value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setCmd).Standalone()

	rootCmd.AddCommand(setCmd)

	carapace.Gen(setCmd).PositionalCompletion(
		carapace.ActionValues(
			"lightness", "hue", "chroma", "lab-a", "lab-b", "oklab-l",
			"oklab-a", "oklab-b", "red", "green", "blue", "hsl-hue",
			"hsl-saturation", "hsl-lightness", "alpha",
		),
	)
}
