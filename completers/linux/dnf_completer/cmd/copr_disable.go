package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var copr_disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "disable specified Copr repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(copr_disableCmd).Standalone()

	coprCmd.AddCommand(copr_disableCmd)
}
