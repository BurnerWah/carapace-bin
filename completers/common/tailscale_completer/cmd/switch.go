package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switch to a different Tailscale account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(switchCmd).Standalone()

	switchCmd.Flags().Bool("list", false, "list available accounts")

	carapace.Gen(switchCmd).PositionalCompletion(carapace.ActionExecCommand("tailscale", "switch", "--list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)

		for _, line := range lines {
			if fields := strings.Fields(line); len(fields) > 1 && fields[0] != "ID" {
				vals = append(vals, fields[0], fields[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}))

	rootCmd.AddCommand(switchCmd)
}
