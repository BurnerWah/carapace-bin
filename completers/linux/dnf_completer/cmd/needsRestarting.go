package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var needsRestartingCmd = &cobra.Command{
	Use:     "needs-restarting",
	Short:   "Determine whether system or systemd services need restarting",
	GroupID: "commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(needsRestartingCmd).Standalone()

	needsRestartingCmd.Flags().BoolP("reboothint", "r", false, "Has no effect, kept for compatibility with DNF 4")
	needsRestartingCmd.Flags().BoolP("services", "s", false, "List systemd services started before their dependencies were updated")
	needsRestartingCmd.Flag("reboothint").Hidden = true
	rootCmd.AddCommand(needsRestartingCmd)
}
