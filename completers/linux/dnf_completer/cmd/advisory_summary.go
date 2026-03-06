package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var advisory_summaryCmd = &cobra.Command{
	Use:     "summary",
	Short:   "Print summary of advisories",
	GroupID: "query",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(advisory_summaryCmd).Standalone()

	advisory_summaryCmd.Flags().String("advisory-severities", "", "Limit to packages in advisories with specified severity")
	advisory_summaryCmd.Flags().Bool("all", false, "Show advisories containing any version of installed packages")
	advisory_summaryCmd.Flags().Bool("available", false, "Show advisories containing newer versions of installed packages")
	advisory_summaryCmd.Flags().Bool("bugfix", false, "Limit to packages in bugfix advisories")
	advisory_summaryCmd.Flags().String("bz", "", "Alias for '--bzs'")
	advisory_summaryCmd.Flags().String("bzs", "", "Limit to packages in advisories that fix a Bugzilla ID")
	advisory_summaryCmd.Flags().String("contains-pkgs", "", "Show only advisories containing packages with specified names")
	advisory_summaryCmd.Flags().String("cve", "", "Alias for '--cves'")
	advisory_summaryCmd.Flags().String("cves", "", "Limit to packages in advisories that fix a CVE")
	advisory_summaryCmd.Flags().Bool("enhancement", false, "Limit to packages in enhancement advisories")
	advisory_summaryCmd.Flags().Bool("installed", false, "Show advisories containing equal and older versions of installed packages")
	advisory_summaryCmd.Flags().Bool("newpackage", false, "Limit to packages in newpackage advisories")
	advisory_summaryCmd.Flags().Bool("security", false, "Limit to packages in security advisories")
	advisory_summaryCmd.Flags().Bool("updates", false, "Show advisories containing newer versions of installed packages for which a newer version is available")
	advisory_summaryCmd.Flags().Bool("with-bz", false, "Show only advisories referencing a bugzilla")
	advisory_summaryCmd.Flags().Bool("with-cve", false, "Show only advisories referencing a CVE")
	advisory_summaryCmd.Flag("bz").Hidden = true
	advisory_summaryCmd.Flag("cve").Hidden = true
	advisoryCmd.AddCommand(advisory_summaryCmd)
}
