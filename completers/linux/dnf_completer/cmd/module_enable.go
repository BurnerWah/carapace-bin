package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var module_enableCmd = &cobra.Command{
	Use:     "enable",
	Short:   "Enable module streams and make their packages available",
	GroupID: "management",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(module_enableCmd).Standalone()

	module_enableCmd.Flags().Bool("skip-broken", false, "Allow resolving of depsolve problems by skipping packages")
	module_enableCmd.Flags().Bool("skip-unavailable", false, "Allow skipping unavailable packages")
	moduleCmd.AddCommand(module_enableCmd)
}
