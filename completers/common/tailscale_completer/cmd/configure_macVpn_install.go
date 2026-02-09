package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configure_macVpn_installCmd = &cobra.Command{
	Use:   "install",
	Short: "Write the Tailscale VPN configuration to the macOS settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configure_macVpn_installCmd).Standalone()

	configure_macVpnCmd.AddCommand(configure_macVpn_installCmd)
}
