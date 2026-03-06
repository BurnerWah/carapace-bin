package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var advisory_infoCmd = &cobra.Command{
	Use:     "info",
	Short:   "Print details about advisories",
	GroupID: "query",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(advisory_infoCmd).Standalone()

	advisory_infoCmd.Flags().String("advisory-severities", "", "Limit to packages in advisories with specified severity")
	advisory_infoCmd.Flags().Bool("all", false, "Show advisories containing any version of installed packages")
	advisory_infoCmd.Flags().Bool("available", false, "Show advisories containing newer versions of installed packages")
	advisory_infoCmd.Flags().Bool("bugfix", false, "Limit to packages in bugfix advisories")
	advisory_infoCmd.Flags().String("bz", "", "Alias for '--bzs'")
	advisory_infoCmd.Flags().String("bzs", "", "Limit to packages in advisories that fix a Bugzilla ID")
	advisory_infoCmd.Flags().String("contains-pkgs", "", "Show only advisories containing packages with specified names")
	advisory_infoCmd.Flags().String("cve", "", "Alias for '--cves'")
	advisory_infoCmd.Flags().String("cves", "", "Limit to packages in advisories that fix a CVE")
	advisory_infoCmd.Flags().Bool("enhancement", false, "Limit to packages in enhancement advisories")
	advisory_infoCmd.Flags().Bool("installed", false, "Show advisories containing equal and older versions of installed packages")
	advisory_infoCmd.Flags().Bool("newpackage", false, "Limit to packages in newpackage advisories")
	advisory_infoCmd.Flags().Bool("security", false, "Limit to packages in security advisories")
	advisory_infoCmd.Flags().Bool("updates", false, "Show advisories containing newer versions of installed packages for which a newer version is available")
	advisory_infoCmd.Flags().Bool("with-bz", false, "Show only advisories referencing a bugzilla")
	advisory_infoCmd.Flags().Bool("with-cve", false, "Show only advisories referencing a CVE")
	advisory_infoCmd.Flag("bz").Hidden = true
	advisory_infoCmd.Flag("cve").Hidden = true
	advisoryCmd.AddCommand(advisory_infoCmd)
}
