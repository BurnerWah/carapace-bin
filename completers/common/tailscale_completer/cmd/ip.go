package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Show Tailscale IP addresses",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ipCmd).Standalone()

	ipCmd.Flags().BoolP("1", "1", false, "only print one IP address")
	ipCmd.Flags().BoolP("4", "4", false, "only print IPv4 address")
	ipCmd.Flags().BoolP("6", "6", false, "only print IPv6 address")

	ipCmd.MarkFlagsMutuallyExclusive("1", "4", "6")

	rootCmd.AddCommand(ipCmd)
}
