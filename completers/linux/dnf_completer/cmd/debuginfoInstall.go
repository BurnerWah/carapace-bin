package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/dnf"
	"github.com/spf13/cobra"
)

var debuginfoInstallCmd = &cobra.Command{
	Use:     "debuginfo-install",
	Short:   "Install debuginfo packages",
	GroupID: "core",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debuginfoInstallCmd).Standalone()

	debuginfoInstallCmd.Flags().Bool("allowerasing", false, "Allow erasing of installed packages to resolve problems")
	debuginfoInstallCmd.Flags().Bool("skip-broken", false, "Allow resolving of depsolve problems by skipping packages")
	debuginfoInstallCmd.Flags().Bool("skip-unavailable", false, "Allow skipping unavailable packages")

	rootCmd.AddCommand(debuginfoInstallCmd)

	carapace.Gen(debuginfoInstallCmd).PositionalAnyCompletion(dnf.ActionPackages(false))
}
