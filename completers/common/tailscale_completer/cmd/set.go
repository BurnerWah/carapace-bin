package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tailscale_completer/cmd/common"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Change specified preferences",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setCmd).Standalone()

	common.AddSettingsFlags(setCmd)

	setCmd.Flags().Bool("auto-update", false, "automatically update to the latest available version")
	setCmd.Flags().String("nickname", "", "nickname for the current account")
	setCmd.Flags().String("relay-server-port", "", "UDP port number (0 will pick a random unused port) for the relay server to bind to, on all interfaces, or empty string to disable relay server functionality")
	// TODO: --relay-server-static-endpoints
	setCmd.Flags().Bool("report-posture", false, "allow management plane to gather device posture information")
	setCmd.Flags().Bool("update-check", false, "notify about available Tailscale updates")
	setCmd.Flags().Bool("webclient", false, "expose the web interface for managing this node over Tailscale at port 5252")

	rootCmd.AddCommand(setCmd)
}
