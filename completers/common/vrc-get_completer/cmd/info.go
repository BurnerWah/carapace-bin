package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Shows information for other program",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()

	infoCmd.Flags().BoolP("help", "h", false, "Print help")
	infoCmd.Flags().BoolP("version", "V", false, "Print version")
	rootCmd.AddCommand(infoCmd)
}
