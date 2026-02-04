package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var info_projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Show project information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(info_projectCmd).Standalone()

	info_projectCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	info_projectCmd.Flags().String("json-format", "", "Output json format")
	info_projectCmd.Flags().StringP("project", "p", "", "Path to project dir. by default CWD or parents of CWD will be used")
	info_projectCmd.Flags().BoolP("version", "V", false, "Print version")
	infoCmd.AddCommand(info_projectCmd)
}
