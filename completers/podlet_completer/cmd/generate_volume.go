package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generate_volumeCmd = &cobra.Command{
	Use:   "volume",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_volumeCmd).Standalone()

	generateCmd.AddCommand(generate_volumeCmd)
}
