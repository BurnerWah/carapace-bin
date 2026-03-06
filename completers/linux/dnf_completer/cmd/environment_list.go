package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var environment_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List environments",
	GroupID: "query",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(environment_listCmd).Standalone()

	environment_listCmd.Flags().Bool("available", false, "Show only available environments")
	environment_listCmd.Flags().Bool("installed", false, "Show only installed environments")
	environmentCmd.AddCommand(environment_listCmd)
}
