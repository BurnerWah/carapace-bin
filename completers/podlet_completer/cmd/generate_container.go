package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generate_containerCmd = &cobra.Command{
	Use:   "container",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_containerCmd).Standalone()

	generateCmd.AddCommand(generate_containerCmd)
}
