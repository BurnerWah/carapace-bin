package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var copr_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove specified Copr repository from the system",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(copr_removeCmd).Standalone()

	coprCmd.AddCommand(copr_removeCmd)
}
