package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/dnf"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install",
	Short:   "Install software",
	GroupID: "core",
	Aliases: []string{"in"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().String("advisories", "", "Limit to packages in advisories with specified name")
	installCmd.Flags().String("advisory", "", "Alias for '--advisories'")
	installCmd.Flags().String("advisory-severities", "", "Limit to packages in advisories with specified severity")
	installCmd.Flags().Bool("allow-downgrade", false, "Allow downgrade of dependencies for resolve of requested operation")
	installCmd.Flags().Bool("allowerasing", false, "Allow erasing of installed packages to resolve problems")
	installCmd.Flags().Bool("bugfix", false, "Limit to packages in bugfix advisories")
	installCmd.Flags().String("bz", "", "Alias for '--bzs'")
	installCmd.Flags().String("bzs", "", "Limit to packages in advisories that fix a Bugzilla ID")
	installCmd.Flags().String("cve", "", "Alias for '--cves'")
	installCmd.Flags().String("cves", "", "Limit to packages in advisories that fix a CVE")
	installCmd.Flags().Bool("downloadonly", false, "Only download packages for a transaction")
	installCmd.Flags().Bool("enhancement", false, "Limit to packages in enhancement advisories")
	installCmd.Flags().Bool("newpackage", false, "Limit to packages in newpackage advisories")
	installCmd.Flags().Bool("no-allow-downgrade", false, "Disable downgrade of dependencies for resolve of requested operation")
	installCmd.Flags().Bool("offline", false, "Store the transaction to be performed offline")
	installCmd.Flags().Bool("security", false, "Limit to packages in security advisories")
	installCmd.Flags().Bool("skip-broken", false, "Allow resolving of depsolve problems by skipping packages")
	installCmd.Flags().Bool("skip-unavailable", false, "Allow skipping unavailable packages")
	installCmd.Flags().String("store", "", "Store the current transaction in a directory instead of running it")
	installCmd.Flag("advisory").Hidden = true
	installCmd.Flag("bz").Hidden = true
	installCmd.Flag("cve").Hidden = true

	installCmd.MarkFlagsMutuallyExclusive("allow-downgrade", "no-allow-downgrade")

	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).PositionalAnyCompletion(carapace.Batch(
		dnf.ActionPackages(true),
		carapace.ActionFiles("rpm"),
	).ToA())
}
