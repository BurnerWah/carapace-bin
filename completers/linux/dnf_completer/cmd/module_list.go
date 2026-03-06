package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var module_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List module streams",
	GroupID: "query",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(module_listCmd).Standalone()

	module_listCmd.Flags().Bool("disabled", false, "Show disabled modules")
	module_listCmd.Flags().Bool("enabled", false, "Show enabled modules")
	moduleCmd.AddCommand(module_listCmd)
}
