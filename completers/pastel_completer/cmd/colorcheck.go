package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var colorcheckCmd = &cobra.Command{
	Use:   "colorcheck",
	Short: "Check if your terminal emulator supports 24-bit colors",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(colorcheckCmd).Standalone()

	rootCmd.AddCommand(colorcheckCmd)
}
