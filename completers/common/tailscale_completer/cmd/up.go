package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tailscale_completer/cmd/common"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Connect to Tailscale, logging in if needed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upCmd).Standalone()

	common.AddSettingsFlags(upCmd)
	common.AddAcceptRisksFlag(upCmd)
	common.AddUpFlags(upCmd)

	upCmd.Flags().Bool("json", false, "output in JSON format")
	upCmd.Flags().Bool("reset", false, "reset unspecified settings to their default values")

	rootCmd.AddCommand(upCmd)
}
