package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var module_disableCmd = &cobra.Command{
	Use:     "disable",
	Short:   "Disable modules including all their streams",
	GroupID: "management",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(module_disableCmd).Standalone()

	module_disableCmd.Flags().Bool("skip-unavailable", false, "Allow skipping unavailable packages")
	moduleCmd.AddCommand(module_disableCmd)
}
