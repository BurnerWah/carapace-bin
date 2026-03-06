package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var markCmd = &cobra.Command{
	Use:     "mark",
	Short:   "Change the reason of an installed package",
	GroupID: "core",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(markCmd).Standalone()

	markCmd.Flags().Bool("skip-unavailable", false, "Allow skipping unavailable packages")
	markCmd.Flags().String("store", "", "Store the current transaction in a directory instead of running it")
	rootCmd.AddCommand(markCmd)
}
