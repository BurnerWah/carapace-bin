package common

import "github.com/spf13/cobra"

func AddSettingsFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("accept-dns", true, "accept DNS configuration from the admin panel")
	cmd.Flags().Bool("accept-routes", true, "accept routes advertised by other Tailscale nodes")
	cmd.Flags().Bool("advertise-connector", false, "advertise this node as an app connector")
	cmd.Flags().Bool("advertise-exit-node", false, "offer to be an exit node for internet traffic for the tailnet")
	// TODO: --advertise-routes
	// TODO: --exit-node
	cmd.Flags().Bool("exit-node-allow-lan-access", false, "Allow direct access to the local network when routing traffic via an exit node")
	cmd.Flags().String("hostname", "", "hostname to use instead of the one provided by the OS")
	cmd.Flags().String("operator", "", "Unix username to allow to operate on tailscaled without sudo") // Add Completions
	cmd.Flags().Bool("shields-up", false, "don't allow incoming connections")
	cmd.Flags().Bool("ssh", false, "run an SSH server, permitting access per tailnet admin's declared policy")
}
