package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/dnf"
	"github.com/spf13/cobra"
)

var downgradeCmd = &cobra.Command{
	Use:     "downgrade",
	Short:   "Downgrade software",
	GroupID: "core",
	Aliases: []string{"dg"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(downgradeCmd).Standalone()

	downgradeCmd.Flags().Bool("allow-downgrade", false, "Allow downgrade of dependencies for resolve of requested operation")
	downgradeCmd.Flags().Bool("allowerasing", false, "Allow erasing of installed packages to resolve problems")
	downgradeCmd.Flags().Bool("downloadonly", false, "Only download packages for a transaction")
	downgradeCmd.Flags().Bool("no-allow-downgrade", false, "Disable downgrade of dependencies for resolve of requested operation")
	downgradeCmd.Flags().Bool("offline", false, "Store the transaction to be performed offline")
	downgradeCmd.Flags().Bool("skip-broken", false, "Allow resolving of depsolve problems by skipping packages")
	downgradeCmd.Flags().Bool("skip-unavailable", false, "Allow skipping unavailable packages")
	downgradeCmd.Flags().String("store", "", "Store the current transaction in a directory instead of running it")

	rootCmd.AddCommand(downgradeCmd)

	carapace.Gen(downgradeCmd).PositionalAnyCompletion(dnf.ActionPackages(false))
}
