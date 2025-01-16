package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mixCmd = &cobra.Command{
	Use:   "mix",
	Short: "Mix two colors in the given colorspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mixCmd).Standalone()

	mixCmd.Flags().StringP("colorspace", "s", "Lab", "The colorspace in which to interpolate")
	mixCmd.Flags().Float64P("fraction", "f", 0.5, "The number between 0.0 and 1.0 determining how much to mix in from the base color")
	rootCmd.AddCommand(mixCmd)

	carapace.Gen(mixCmd).FlagCompletion(carapace.ActionMap{
		"colorspace": carapace.ActionValues("Lab", "LCh", "RGB", "HSL", "OkLab"),
	})
}
