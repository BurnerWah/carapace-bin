package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generate_imageCmd = &cobra.Command{
	Use:   "image",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_imageCmd).Standalone()

	generateCmd.AddCommand(generate_imageCmd)
}
