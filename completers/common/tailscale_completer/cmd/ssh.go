package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var sshCmd = &cobra.Command{
	Use:                "ssh",
	Short:              "SSH to a Tailscale machine",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(sshCmd).Standalone()

	rootCmd.AddCommand(sshCmd)

	// This is just an SSH wrapper apparently
	carapace.Gen(sshCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("ssh"),
	)
}
