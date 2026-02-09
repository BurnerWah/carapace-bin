package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configure_macVpn_uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Delete the Tailscale VPN configuration from the macOS settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configure_macVpn_uninstallCmd).Standalone()

	configure_macVpnCmd.AddCommand(configure_macVpn_uninstallCmd)
}
