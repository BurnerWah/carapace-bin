package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pipeCmd = &cobra.Command{
	Use:   "pipe",
	Short: "Output metrics in JSON format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pipeCmd).Standalone()

	pipeCmd.Flags().IntP("samples", "s", 0, "Number of samples to run for")
	rootCmd.AddCommand(pipeCmd)
}
