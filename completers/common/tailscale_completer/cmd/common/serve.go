package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

func AddServeFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("bg", false, "Run the command as a background process")
	cmd.Flags().String("set-path", "", "Appends the specified path to the base URL for accessing the underlying service")
	cmd.Flags().Uint("https", 0, "Expose an HTTPS server at the specified port (default mode)")
	cmd.Flags().Uint("tcp", 0, "Expose a TCP forwarder to forward raw TCP packets at the specified port")
	cmd.Flags().Uint("tls-terminated-tcp", 0, "Expose a TCP forwarder to forward TLS-terminated TCP packets at the specified port")
	cmd.Flags().Uint("proxy-protocol", 0, "PROXY protocol version (1 or 2) for TCP forwarding")
	cmd.Flags().Bool("yes", false, "Update without interactive prompts")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"https":              net.ActionPorts(),
		"tcp":                net.ActionPorts(),
		"tls-terminated-tcp": net.ActionPorts(),
		"proxy-protocol":     carapace.ActionValues("1", "2"),
	})
}
