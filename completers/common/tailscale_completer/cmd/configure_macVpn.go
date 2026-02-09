package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configure_macVpnCmd = &cobra.Command{
	Use:   "macVpn",
	Short: "Manage the VPN configuration on macOS",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configure_macVpnCmd).Standalone()

	configureCmd.AddCommand(configure_macVpnCmd)
}
