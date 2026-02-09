package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configure_sysext_activateCmd = &cobra.Command{
	Use:   "activate",
	Short: "Register the Tailscale system extension with macOS",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configure_sysext_activateCmd).Standalone()

	configure_sysextCmd.AddCommand(configure_sysext_activateCmd)
}
