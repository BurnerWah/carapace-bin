package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dns_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Print the current DNS status and configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dns_statusCmd).Standalone()

	dns_statusCmd.Flags().Bool("all", false, "outputs advanced debugging information")

	dnsCmd.AddCommand(dns_statusCmd)
}
