package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lock_revokeKeysCmd = &cobra.Command{
	Use:   "revoke-keys",
	Short: "Revoke compromised tailnet-lock keys",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lock_revokeKeysCmd).Standalone()

	lock_revokeKeysCmd.Flags().Bool("cosign", false, "continue generating the recovery using the tailnet lock key on this device and the provided recovery blob")
	lock_revokeKeysCmd.Flags().Bool("finish", false, "finish the recovery process by transmitting the revocation")
	lock_revokeKeysCmd.Flags().String("fork-from", "", "parent AUM hash to rewrite from (advanced users only)")

	lockCmd.AddCommand(lock_revokeKeysCmd)

	// TODO: Add positional completion
}
