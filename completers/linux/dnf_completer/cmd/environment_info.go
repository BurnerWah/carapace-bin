package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var environment_infoCmd = &cobra.Command{
	Use:     "info",
	Short:   "Print details about environments",
	GroupID: "query",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(environment_infoCmd).Standalone()

	environment_infoCmd.Flags().Bool("available", false, "Show only available environments")
	environment_infoCmd.Flags().Bool("installed", false, "Show only installed environments")
	environmentCmd.AddCommand(environment_infoCmd)
}
