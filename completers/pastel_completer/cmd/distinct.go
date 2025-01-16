package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var distinctCmd = &cobra.Command{
	Use:   "distinct",
	Short: "Generate a set of visually distinct colors",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(distinctCmd).Standalone()

	distinctCmd.Flags().StringP("metric", "m", "CIE76", "Distance metric to compute mutual color distances")
	distinctCmd.Flags().BoolP("verbose", "v", false, "Print simulation output to STDERR")
	rootCmd.AddCommand(distinctCmd)

	carapace.Gen(distinctCmd).FlagCompletion(carapace.ActionMap{
		"metric": carapace.ActionValues("CIEDE2000", "CIE76"),
	})
}
