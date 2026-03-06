package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repoclosureCmd = &cobra.Command{
	Use:     "repoclosure",
	Short:   "Print list of unresolved dependencies for repositories",
	GroupID: "commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repoclosureCmd).Standalone()

	repoclosureCmd.Flags().String("arch", "", "Only check packages of specified architectures")
	repoclosureCmd.Flags().String("check", "", "Specify repo ids to check")
	repoclosureCmd.Flags().Bool("newest", false, "Only consider the latest version of a package from each repo")
	rootCmd.AddCommand(repoclosureCmd)
}
