package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configure_sysext_deactivateCmd = &cobra.Command{
	Use:   "deactivate",
	Short: "Deactivate the Tailscale system extension on macOS",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configure_sysext_deactivateCmd).Standalone()

	configure_sysextCmd.AddCommand(configure_sysext_deactivateCmd)
}
