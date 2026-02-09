package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func AddUpFlags(cmd *cobra.Command) {
	// TODO: --advertise-tags
	cmd.Flags().String("audience", "", "Audience used when requesting an ID token from an identity provider for auth keys via workload identity federation")
	cmd.Flags().String("auth-key", "", "node authorization key")
	cmd.Flags().String("client-id", "", "Client ID used to generate authkeys via workload identity federation")
	cmd.Flags().String("client-secret", "", "Client Secret used to generate authkeys via OAuth")
	cmd.Flags().Bool("force-reauth", false, "force reauthentication")
	cmd.Flags().String("id-token", "", "ID token from the identity provider to exchange with the control server for workload identity federation")
	cmd.Flags().String("login-server", "https://controlplane.tailscale.com", "base URL of control server")
	cmd.Flags().Bool("qr", false, "show QR code for login URLs")
	cmd.Flags().String("qr-format", "auto", "QR code formatting")
	// TODO: Add linux-only flags
	cmd.Flags().Duration("timeout", 0, "maximum amount of time to wait for tailscaled to enter a Running state; default (0s) blocks forever")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"qr-format": carapace.ActionValues("auto", "ascii", "large", "small"),
	})
}
