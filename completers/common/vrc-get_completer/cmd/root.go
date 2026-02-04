package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vrc-get",
	Short: "Open Source command line client of VRChat Package Manager",
	Long:  "https://github.com/vrc-get/vrc-get",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")
}
