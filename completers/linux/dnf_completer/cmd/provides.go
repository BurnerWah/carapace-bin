package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var providesCmd = &cobra.Command{
	Use:     "provides",
	Short:   "Find what package provides the given value",
	GroupID: "core",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(providesCmd).Standalone()

	rootCmd.AddCommand(providesCmd)
}
