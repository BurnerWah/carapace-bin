package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var advisory_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List advisories",
	GroupID: "query",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(advisory_listCmd).Standalone()

	advisory_listCmd.Flags().String("advisory-severities", "", "Limit to packages in advisories with specified severity")
	advisory_listCmd.Flags().Bool("all", false, "Show advisories containing any version of installed packages")
	advisory_listCmd.Flags().Bool("available", false, "Show advisories containing newer versions of installed packages")
	advisory_listCmd.Flags().Bool("bugfix", false, "Limit to packages in bugfix advisories")
	advisory_listCmd.Flags().String("bz", "", "Alias for '--bzs'")
	advisory_listCmd.Flags().String("bzs", "", "Limit to packages in advisories that fix a Bugzilla ID")
	advisory_listCmd.Flags().String("contains-pkgs", "", "Show only advisories containing packages with specified names")
	advisory_listCmd.Flags().String("cve", "", "Alias for '--cves'")
	advisory_listCmd.Flags().String("cves", "", "Limit to packages in advisories that fix a CVE")
	advisory_listCmd.Flags().Bool("enhancement", false, "Limit to packages in enhancement advisories")
	advisory_listCmd.Flags().Bool("installed", false, "Show advisories containing equal and older versions of installed packages")
	advisory_listCmd.Flags().Bool("json", false, "Request json output format")
	advisory_listCmd.Flags().Bool("newpackage", false, "Limit to packages in newpackage advisories")
	advisory_listCmd.Flags().Bool("security", false, "Limit to packages in security advisories")
	advisory_listCmd.Flags().Bool("updates", false, "Show advisories containing newer versions of installed packages for which a newer version is available")
	advisory_listCmd.Flags().Bool("with-bz", false, "Show only advisories referencing a bugzilla")
	advisory_listCmd.Flags().Bool("with-cve", false, "Show only advisories referencing a CVE")
	advisory_listCmd.Flag("bz").Hidden = true
	advisory_listCmd.Flag("cve").Hidden = true
	advisoryCmd.AddCommand(advisory_listCmd)
}
