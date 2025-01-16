package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var formatCmd = &cobra.Command{
	Use:   "format",
	Short: "Convert a color to the given format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(formatCmd).Standalone()

	rootCmd.AddCommand(formatCmd)

	carapace.Gen(formatCmd).PositionalCompletion(
		carapace.ActionValues(
			"rgb", "rgb-float", "hex", "hsl", "hsl-hue", "hsl-saturation",
			"hsl-lightness", "hsv", "hsv-hue", "hsv-saturation", "hsv-value",
			"lch", "lch-lightness", "lch-chroma", "lch-hue", "lab", "lab-a",
			"lab-b", "oklab", "oklab-l", "oklab-a", "oklab-b", "luminance",
			"brightness", "ansi-8bit", "ansi-24bit", "ansi-8bit-escapecode",
			"ansi-24bit-escapecode", "cmyk", "name",
		),
	)
}
