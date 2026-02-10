package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appcRoutesCmd = &cobra.Command{
	Use:   "appc-routes",
	Short: "Print the current app connector routes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appcRoutesCmd).Standalone()

	appcRoutesCmd.Flags().Bool("all", false, "Print learned domains and routes and extra policy configured routes")
	appcRoutesCmd.Flags().Bool("map", false, "Print the map of learned domains: [routes]")
	appcRoutesCmd.Flags().Bool("n", false, "Print the total number of routes this node advertises")

	rootCmd.AddCommand(appcRoutesCmd)
}
