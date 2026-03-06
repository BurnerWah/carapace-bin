package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var copr_listCmd = &cobra.Command{
	Use:   "list",
	Short: "list Copr repositories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(copr_listCmd).Standalone()

	coprCmd.AddCommand(copr_listCmd)
}
