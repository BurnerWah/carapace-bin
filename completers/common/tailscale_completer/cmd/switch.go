package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switch to a different Tailscale account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(switchCmd).Standalone()

	switchCmd.Flags().Bool("list", false, "list available accounts")

	rootCmd.AddCommand(switchCmd)
}
