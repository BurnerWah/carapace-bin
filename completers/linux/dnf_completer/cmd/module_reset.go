package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var module_resetCmd = &cobra.Command{
	Use:     "reset",
	Short:   "Reset module state so it's no longer enabled or disabled",
	GroupID: "management",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(module_resetCmd).Standalone()

	module_resetCmd.Flags().Bool("skip-unavailable", false, "Allow skipping unavailable packages")
	moduleCmd.AddCommand(module_resetCmd)
}
