package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pastel",
	Short: "A command-line tool to generate, analyze, convert and manipulate colors",
	Long:  "https://github.com/sharkdp/pastel",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("color-mode", "m", "", "Specify the terminal color mode")
	rootCmd.Flags().String("color-picker", "", "Use a specific tool to pick the colors")
	rootCmd.Flags().BoolP("force-color", "f", false, "Alias for --mode=24bit")
	rootCmd.Flags().BoolP("version", "V", false, "Print version information")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Print help information")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"color-mode":   carapace.ActionValues("24bit", "8bit", "off", "auto"),
		"color-picker": carapace.ActionValues("osascript", "gpick", "xcolor", "wcolor, grabc, colorpicker, chameleon, kcolorchooser, zenity, yad, hyprpicker"),
	})
}
