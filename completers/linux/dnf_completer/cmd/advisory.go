package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var advisoryCmd = &cobra.Command{
	Use:     "advisory",
	Short:   "Manage advisories",
	GroupID: "subcommands",
	Aliases: []string{"updateinfo"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(advisoryCmd).Standalone()
	advisoryCmd.AddGroup(
		&cobra.Group{ID: "query", Title: ""},
	)

	rootCmd.AddCommand(advisoryCmd)
}
