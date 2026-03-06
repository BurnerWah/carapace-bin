package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/dnf"
	"github.com/spf13/cobra"
)

var checkUpgradeCmd = &cobra.Command{
	Use:     "check-upgrade",
	Short:   "Check for available package upgrades",
	GroupID: "core",
	Aliases: []string{"check-update"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkUpgradeCmd).Standalone()

	checkUpgradeCmd.Flags().String("advisories", "", "Limit to packages in advisories with specified name")
	checkUpgradeCmd.Flags().String("advisory", "", "Alias for '--advisories'")
	checkUpgradeCmd.Flags().String("advisory-severities", "", "Limit to packages in advisories with specified severity")
	checkUpgradeCmd.Flags().Bool("bugfix", false, "Limit to packages in bugfix advisories")
	checkUpgradeCmd.Flags().String("bz", "", "Alias for '--bzs'")
	checkUpgradeCmd.Flags().String("bzs", "", "Limit to packages in advisories that fix a Bugzilla ID")
	checkUpgradeCmd.Flags().Bool("changelogs", false, "Show changelogs before update.")
	checkUpgradeCmd.Flags().String("cve", "", "Alias for '--cves'")
	checkUpgradeCmd.Flags().String("cves", "", "Limit to packages in advisories that fix a CVE")
	checkUpgradeCmd.Flags().Bool("enhancement", false, "Limit to packages in enhancement advisories")
	checkUpgradeCmd.Flags().Bool("minimal", false, "Reports the lowest versions of packages that fix advisories of type bugfix, enhancement, security, or newpackage")
	checkUpgradeCmd.Flags().Bool("newpackage", false, "Limit to packages in newpackage advisories")
	checkUpgradeCmd.Flags().Bool("security", false, "Limit to packages in security advisories")
	checkUpgradeCmd.Flag("advisory").Hidden = true
	checkUpgradeCmd.Flag("bz").Hidden = true
	checkUpgradeCmd.Flag("cve").Hidden = true

	rootCmd.AddCommand(checkUpgradeCmd)

	carapace.Gen(checkUpgradeCmd).PositionalAnyCompletion(dnf.ActionPackages(false))
}
