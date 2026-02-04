package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update local repository cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()

	updateCmd.Flags().BoolP("help", "h", false, "Print help")
	updateCmd.Flags().BoolP("version", "V", false, "Print version")
	rootCmd.AddCommand(updateCmd)
}
