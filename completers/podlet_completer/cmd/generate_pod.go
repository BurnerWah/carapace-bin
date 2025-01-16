package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generate_podCmd = &cobra.Command{
	Use:   "pod",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_podCmd).Standalone()

	generateCmd.AddCommand(generate_podCmd)
}
