package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tailscale_completer/cmd/common"
	"github.com/spf13/cobra"
)

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Disconnect from Tailscale",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(downCmd).Standalone()

	common.AddAcceptRisksFlag(downCmd)

	downCmd.Flags().String("reason", "", "reason for the disconnect, if required by a policy")

	rootCmd.AddCommand(downCmd)
}
