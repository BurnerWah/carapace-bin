package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tailscale_completer/cmd/common"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to a Tailscale account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loginCmd).Standalone()

	common.AddSettingsFlags(loginCmd)
	common.AddAcceptRisksFlag(loginCmd)
	common.AddUpFlags(loginCmd)

	loginCmd.Flags().String("nickname", "", "nickname for the current account")

	rootCmd.AddCommand(loginCmd)
}
