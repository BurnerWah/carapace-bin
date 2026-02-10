package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tailscale_completer/cmd/common"
	"github.com/spf13/cobra"
)

var funnelCmd = &cobra.Command{
	Use:   "funnel",
	Short: "Serve content and local servers on the internet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(funnelCmd).Standalone()

	common.AddServeFlags(funnelCmd)

	rootCmd.AddCommand(funnelCmd)

	// TODO: Add positional completions
}
