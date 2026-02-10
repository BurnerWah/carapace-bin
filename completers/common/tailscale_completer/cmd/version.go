package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print Tailscale version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCmd).Standalone()

	versionCmd.Flags().Bool("daemon", false, "also print local node's daemon version")
	versionCmd.Flags().Bool("json", false, "output in JSON format")
	versionCmd.Flags().Bool("upstream", false, "fetch and print the latest upstream release version from pkgs.tailscale.com")

	rootCmd.AddCommand(versionCmd)
}
