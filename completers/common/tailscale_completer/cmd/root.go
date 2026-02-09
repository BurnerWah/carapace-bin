package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tailscale",
	Short: "The easiest, most secure way to use WireGuard",
	Long:  "https://tailscale.com/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("socket", "", "path to tailscaled socket")
	rootCmd.Flags().BoolP("version", "V", false, "Print Tailscale version")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Print help")

	rootCmd.Flag("version").Hidden = true

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"socket": carapace.ActionFiles(),
	})
}
