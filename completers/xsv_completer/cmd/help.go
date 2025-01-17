package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Show help",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(helpCmd).Standalone()

	helpCmd.Flags().BoolP("help", "h", false, "Display this message")
	helpCmd.Flags().Bool("list", false, "List all commands available.")
	helpCmd.Flags().Bool("version", false, "Print version info and exit")
	rootCmd.AddCommand(helpCmd)
}
