package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func AddDelimiterFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("delimiter", "d", "", "The field delimiter for reading CSV data")
}

func AddNoHeadersFlag(cmd *cobra.Command) {
	cmd.Flags().BoolP("no-headers", "n", false, "Don't interpert the first row as a header")
}

func AddOutputFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("output", "o", "", "Write output to a file instead of stdout")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
	})
}
