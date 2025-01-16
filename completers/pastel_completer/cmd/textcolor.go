package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var textcolorCmd = &cobra.Command{
	Use:   "textcolor",
	Short: "Get a readable text color for the given background color",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(textcolorCmd).Standalone()

	rootCmd.AddCommand(textcolorCmd)
}
