package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var colorblindCmd = &cobra.Command{
	Use:   "colorblind",
	Short: "Simulate a color under a certain colorblindness profile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(colorblindCmd).Standalone()

	rootCmd.AddCommand(colorblindCmd)

	carapace.Gen(colorblindCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"prot", "protanopia",
			"deuter", "deuteranopia",
			"trit", "tritanopia",
		),
	)
}
