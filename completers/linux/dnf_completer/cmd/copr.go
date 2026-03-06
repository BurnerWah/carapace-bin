package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var coprCmd = &cobra.Command{
	Use:     "copr",
	Short:   "Manage Copr repositories (add-ons provided by users/community/third-party)",
	GroupID: "commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(coprCmd).Standalone()

	coprCmd.Flags().String("hub", "", "Copr hub (the web-UI/API server) hostname")
	rootCmd.AddCommand(coprCmd)
}
