package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completionCmd).Standalone()

	completionCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(completionCmd)
}
