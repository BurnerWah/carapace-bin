package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var podman_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(podman_helpCmd).Standalone()


	podmanCmd.AddCommand(podman_helpCmd)
}
