package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generate_networkCmd = &cobra.Command{
	Use:   "network",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_networkCmd).Standalone()

	generateCmd.AddCommand(generate_networkCmd)
}
