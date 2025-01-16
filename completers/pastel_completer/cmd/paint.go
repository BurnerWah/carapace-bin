package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paintCmd = &cobra.Command{
	Use:   "paint",
	Short: "Print colored text using ANSI escape sequences",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paintCmd).Standalone()

	paintCmd.Flags().BoolP("bold", "b", false, "Print the text in bold face")
	paintCmd.Flags().BoolP("italic", "i", false, "Print the text in italic font")
	paintCmd.Flags().BoolP("no-newline", "n", false, "Do not print a trailing newline character")
	paintCmd.Flags().StringP("on", "o", "", "Use the specified background color")
	paintCmd.Flags().BoolP("underline", "u", false, "Draw a line below the text")
	rootCmd.AddCommand(paintCmd)
}
