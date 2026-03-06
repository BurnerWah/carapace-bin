package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makecacheCmd = &cobra.Command{
	Use:     "makecache",
	Short:   "Generate the metadata cache",
	GroupID: "commands",
	Aliases: []string{"mc"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makecacheCmd).Standalone()

	rootCmd.AddCommand(makecacheCmd)
}
