package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gradientCmd = &cobra.Command{
	Use:   "gradient",
	Short: "Generate an interpolating sequence of colors",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gradientCmd).Standalone()

	gradientCmd.Flags().StringP("colorspace", "s", "Lab", "The colorspace in which to interpolate")
	gradientCmd.Flags().IntP("number", "n", 10, "Number of colors to generate")
	rootCmd.AddCommand(gradientCmd)

	carapace.Gen(gradientCmd).FlagCompletion(carapace.ActionMap{
		"colorspace": carapace.ActionValues("Lab", "LCh", "RGB", "HSL", "OkLab"),
	})
}
