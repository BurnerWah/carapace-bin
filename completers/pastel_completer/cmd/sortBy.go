package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sortByCmd = &cobra.Command{
	Use:   "sort-by",
	Short: "Sort colors by the given property",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sortByCmd).Standalone()

	sortByCmd.Flags().BoolP("reverse", "r", false, "Reverse the sort order")
	sortByCmd.Flags().BoolP("unique", "u", false, "Remove duplicate colors (equality is determined via RGB values)")
	rootCmd.AddCommand(sortByCmd)

	carapace.Gen(sortByCmd).PositionalCompletion(
		carapace.ActionValues("brightness", "luminance", "hue", "chroma", "random"),
	)
}
