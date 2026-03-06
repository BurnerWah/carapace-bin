package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/dnf"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:     "upgrade",
	Short:   "Upgrade software",
	GroupID: "core",
	Aliases: []string{"up", "update"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgradeCmd).Standalone()

	upgradeCmd.Flags().String("advisories", "", "Limit to packages in advisories with specified name")
	upgradeCmd.Flags().String("advisory", "", "Alias for '--advisories'")
	upgradeCmd.Flags().String("advisory-severities", "", "Limit to packages in advisories with specified severity")
	upgradeCmd.Flags().Bool("allow-downgrade", false, "Allow downgrade of dependencies for resolve of requested operation")
	upgradeCmd.Flags().Bool("allowerasing", false, "Allow erasing of installed packages to resolve problems")
	upgradeCmd.Flags().Bool("bugfix", false, "Limit to packages in bugfix advisories")
	upgradeCmd.Flags().String("bz", "", "Alias for '--bzs'")
	upgradeCmd.Flags().String("bzs", "", "Limit to packages in advisories that fix a Bugzilla ID")
	upgradeCmd.Flags().String("cve", "", "Alias for '--cves'")
	upgradeCmd.Flags().String("cves", "", "Limit to packages in advisories that fix a CVE")
	upgradeCmd.Flags().String("destdir", "", "Set directory used for downloading packages to")
	upgradeCmd.Flags().Bool("downloadonly", false, "Only download packages for a transaction")
	upgradeCmd.Flags().Bool("enhancement", false, "Limit to packages in enhancement advisories")
	upgradeCmd.Flags().Bool("minimal", false, "Upgrade packages only to the lowest versions that fix advisories of type bugfix, enhancement, security, or newpackage")
	upgradeCmd.Flags().Bool("newpackage", false, "Limit to packages in newpackage advisories")
	upgradeCmd.Flags().Bool("no-allow-downgrade", false, "Disable downgrade of dependencies for resolve of requested operation")
	upgradeCmd.Flags().Bool("offline", false, "Store the transaction to be performed offline")
	upgradeCmd.Flags().Bool("security", false, "Limit to packages in security advisories")
	upgradeCmd.Flags().Bool("skip-unavailable", false, "Allow skipping unavailable packages")
	upgradeCmd.Flags().String("store", "", "Store the current transaction in a directory instead of running it")
	upgradeCmd.Flag("advisory").Hidden = true
	upgradeCmd.Flag("bz").Hidden = true
	upgradeCmd.Flag("cve").Hidden = true

	rootCmd.AddCommand(upgradeCmd)

	carapace.Gen(upgradeCmd).PositionalAnyCompletion(dnf.ActionPackages(false))
}
