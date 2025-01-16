package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightenCmd = &cobra.Command{
	Use:   "lighten",
	Short: "Lighten color by a specified amount",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightenCmd).Standalone()

	rootCmd.AddCommand(lightenCmd)
}
