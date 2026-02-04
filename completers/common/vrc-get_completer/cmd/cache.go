package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Commands about cache control",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cacheCmd).Standalone()

	cacheCmd.Flags().BoolP("help", "h", false, "Print help")
	cacheCmd.Flags().BoolP("version", "V", false, "Print version")
	rootCmd.AddCommand(cacheCmd)
}
