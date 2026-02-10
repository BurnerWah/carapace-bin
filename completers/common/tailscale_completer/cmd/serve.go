package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tailscale_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve content and local servers on your tailnet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serveCmd).Standalone()

	common.AddServeFlags(serveCmd)

	serveCmd.Flags().String("accept-app-caps", "", "App capabilities to forward to the server")
	serveCmd.Flags().Uint("http", 0, "Expose an HTTP server at the specified port")
	serveCmd.Flags().String("service", "", "Serve for a service with distinct virtual IP instead on node itself")
	serveCmd.Flags().Bool("tun", false, "Forward all traffic to the local machine")

	rootCmd.AddCommand(serveCmd)

	// TODO: Add positional completion

	carapace.Gen(serveCmd).FlagCompletion(carapace.ActionMap{
		"http": net.ActionPorts(),
	})
}
