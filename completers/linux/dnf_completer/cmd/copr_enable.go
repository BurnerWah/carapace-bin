package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var copr_enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "download & install a repository info from a Copr server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(copr_enableCmd).Standalone()

	coprCmd.AddCommand(copr_enableCmd)
}
