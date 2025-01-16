package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func AddFilesystemFlags(cmd *cobra.Command) {
	cmd.Flags().String("files-ignore-unknown", "", "Ignore unknown files")
	cmd.Flags().Int("files-max-size", 0, "Maximum source file size")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"files-ignore-unknown": carapace.ActionValues("true", "false"),
	})
}
