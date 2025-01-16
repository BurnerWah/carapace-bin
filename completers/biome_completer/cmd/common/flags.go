package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func AddConfigPathFlag(cmd *cobra.Command) {
	cmd.Flags().String("config-path", "", "Set the path of the config file or directory")
	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"config-path": carapace.Batch(
			carapace.ActionFiles(".json", ".jsonc"),
			carapace.ActionDirectories(),
		).ToA(),
	})
}

func AddStdinFilePathFlag(cmd *cobra.Command) {
	cmd.Flags().String("stdin-file-path", "", "Path to assume stdin comes from")
	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"stdin-file-path": carapace.ActionFiles(),
	})
}
