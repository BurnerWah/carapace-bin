package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var dns_queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Perform a DNS query",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dns_queryCmd).Standalone()

	dnsCmd.AddCommand(dns_queryCmd)

	carapace.Gen(dns_queryCmd).PositionalCompletion(
		net.ActionHosts(), // I don't like this
		carapace.ActionValues("a", "aaaa", "cname", "mx", "ns", "opt", "ptr", "srv", "txt"),
	)
}
