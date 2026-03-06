package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var replayCmd = &cobra.Command{
	Use:     "replay",
	Short:   "Replay a transaction that was previously stored to a directory",
	GroupID: "core",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(replayCmd).Standalone()

	replayCmd.Flags().Bool("ignore-extras", false, "Don't consider extra packages pulled into the transaction as errors")
	replayCmd.Flags().Bool("ignore-installed", false, "Don't consider mismatches between installed and stored transaction packages as errors")
	replayCmd.Flags().Bool("skip-broken", false, "Allow resolving of depsolve problems by skipping packages")
	replayCmd.Flags().Bool("skip-unavailable", false, "Allow skipping unavailable packages")
	rootCmd.AddCommand(replayCmd)
}
