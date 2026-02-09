package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func AddAcceptRisksFlag(cmd *cobra.Command) {
	cmd.Flags().StringSlice("accept-risk", nil, "accept risk and skip confirmation for risks")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"accept-risk": carapace.ActionValues(
			"lose-ssh", "mac-app-connector", "linux-strict-rp-filter", "all",
		).UniqueList(","),
	})
}
