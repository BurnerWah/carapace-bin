package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/biome_completer/cmd/common"
	"github.com/spf13/cobra"
)

var lspProxyCmd = &cobra.Command{
	Use:   "lsp-proxy",
	Short: "Acts as a server for the Language Server Protocol over stdin/stdout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lspProxyCmd).Standalone()

	common.AddServerFlags(lspProxyCmd)
	rootCmd.AddCommand(lspProxyCmd)
}
