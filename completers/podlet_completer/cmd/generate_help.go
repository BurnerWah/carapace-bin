package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generate_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_helpCmd).Standalone()


	generateCmd.AddCommand(generate_helpCmd)
}
