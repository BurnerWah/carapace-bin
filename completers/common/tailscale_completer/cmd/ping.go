package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tailscale_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping a host at the Tailscale layer, see how it routed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pingCmd).Standalone()

	pingCmd.Flags().Int("c", 10, "max number of pings to send (0 for infinity)")
	pingCmd.Flags().Bool("icmp", false, "do a ICMP-level ping")
	pingCmd.Flags().Bool("peerapi", false, "try hitting the peer's peerapi HTTP server")
	pingCmd.Flags().Int("size", 0, "size of the ping message")
	pingCmd.Flags().Duration("timeout", 5, "timeout before giving up on a ping") // default: 5*time.Second
	pingCmd.Flags().Bool("tsmp", false, "do a TSMP-level ping")
	pingCmd.Flags().Bool("until-direct", true, "stop once a direct path is established")
	pingCmd.Flags().Bool("verbose", false, "verbose output")

	rootCmd.AddCommand(pingCmd)

	// TODO: Complete more than just tailscale peers
	carapace.Gen(pingCmd).PositionalCompletion(action.ActionPeers())
}
