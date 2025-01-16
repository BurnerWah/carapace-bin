package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var darkenCmd = &cobra.Command{
	Use:   "darken",
	Short: "Darken color by a specified amount",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(darkenCmd).Standalone()

	rootCmd.AddCommand(darkenCmd)
}
