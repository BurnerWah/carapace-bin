package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var module_infoCmd = &cobra.Command{
	Use:     "info",
	Short:   "Print module information",
	GroupID: "query",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(module_infoCmd).Standalone()

	module_infoCmd.Flags().Bool("disabled", false, "Show disabled modules")
	module_infoCmd.Flags().Bool("enabled", false, "Show enabled modules")
	moduleCmd.AddCommand(module_infoCmd)
}
