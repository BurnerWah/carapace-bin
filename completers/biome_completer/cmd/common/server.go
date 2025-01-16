package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func AddServerFlags(cmd *cobra.Command) {
	cmd.Flags().String("log-path", "", "Directory to store logs in")
	cmd.Flags().String("log-prefix-name", "server.log", "Prefix applied to log file names")

	AddConfigPathFlag(cmd)

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"log-path": carapace.ActionDirectories(),
	})
}
