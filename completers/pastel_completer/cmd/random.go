package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Generate a list of random colors",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(randomCmd).Standalone()

	randomCmd.Flags().StringP("number", "n", "", "Number of colors to generate")
	randomCmd.Flags().StringP("strategy", "s", "", "Randomization strategy")
	rootCmd.AddCommand(randomCmd)

	carapace.Gen(randomCmd).FlagCompletion(carapace.ActionMap{
		"strategy": carapace.ActionValuesDescribed(
			"vivid", "random hue, limited saturation and lightness values",
			"rgb", "samples uniformly in RGB space",
			"gray", "random gray tone (uniform)",
			"lch_hue", "random hue, fixed lightness and chroma",
		),
	})
}
