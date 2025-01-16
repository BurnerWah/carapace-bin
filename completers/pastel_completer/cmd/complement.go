package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var complementCmd = &cobra.Command{
	Use:   "complement",
	Short: "Get the complementary color (hue rotated by 180°)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(complementCmd).Standalone()

	rootCmd.AddCommand(complementCmd)
}
