package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "macmon",
	Short: "Sudoless performance monitoring CLI tool for Apple Silicon processors",
	Long:  "https://github.com/vladkens/macmon",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().BoolP("help", "h", false, "Print help")
	rootCmd.PersistentFlags().IntP("interval", "i", 1000, "Update interval in milliseconds")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")
}
