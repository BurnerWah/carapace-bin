package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configure_sysext_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Print the enablement status of the Tailscale system extension",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configure_sysext_statusCmd).Standalone()

	configure_sysextCmd.AddCommand(configure_sysext_statusCmd)
}
