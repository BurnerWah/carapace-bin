package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repoqueryCmd = &cobra.Command{
	Use:     "repoquery",
	Short:   "Search for packages matching various criteria",
	GroupID: "query",
	Aliases: []string{"rq"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repoqueryCmd).Standalone()

	repoqueryCmd.Flags().String("advisories", "", "Limit to packages in advisories with specified name")
	repoqueryCmd.Flags().String("advisory", "", "Alias for '--advisories'")
	repoqueryCmd.Flags().String("advisory-severities", "", "Limit to packages in advisories with specified severity")
	repoqueryCmd.Flags().String("arch", "", "Limit to packages of these architectures")
	repoqueryCmd.Flags().Bool("available", false, "Query available packages (default)")
	repoqueryCmd.Flags().Bool("bugfix", false, "Limit to packages in bugfix advisories")
	repoqueryCmd.Flags().String("bz", "", "Alias for '--bzs'")
	repoqueryCmd.Flags().String("bzs", "", "Limit to packages in advisories that fix a Bugzilla ID")
	repoqueryCmd.Flags().Bool("changelogs", false, "Display package changelogs")
	repoqueryCmd.Flags().Bool("conflicts", false, "Like --queryformat=\"%{conflicts}\" but deduplicated and sorted")
	repoqueryCmd.Flags().String("cve", "", "Alias for '--cves'")
	repoqueryCmd.Flags().String("cves", "", "Limit to packages in advisories that fix a CVE")
	repoqueryCmd.Flags().Bool("depends", false, "Like --queryformat=\"%{depends}\" but deduplicated and sorted")
	repoqueryCmd.Flags().Bool("disable-modular-filtering", false, "Include packages of inactive module streams")
	repoqueryCmd.Flags().Bool("duplicates", false, "Limit to installed duplicate packages")
	repoqueryCmd.Flags().Bool("enhancement", false, "Limit to packages in enhancement advisories")
	repoqueryCmd.Flags().Bool("enhances", false, "Like --queryformat=\"%{enhances}\" but deduplicated and sorted")
	repoqueryCmd.Flags().Bool("exactdeps", false, "Limit to packages that require <capability> specified by --whatrequires")
	repoqueryCmd.Flags().Bool("extras", false, "Limit to installed packages that are not present in any available repository")
	repoqueryCmd.Flags().String("file", "", "Limit to packages that own these files")
	repoqueryCmd.Flags().Bool("files", false, "Like --queryformat=\"%{files}\" but deduplicated and sorted")
	repoqueryCmd.Flags().Bool("info", false, "Show detailed information about the packages")
	repoqueryCmd.Flags().Bool("installed", false, "Query installed packages")
	repoqueryCmd.Flags().Bool("installonly", false, "Limit to installed installonly packages")
	repoqueryCmd.Flags().String("latest-limit", "", "Limit to N latest packages for a given name.arch (or all except N latest if N is negative).")
	repoqueryCmd.Flags().Bool("leaves", false, "Limit to groups of installed packages not required by other installed packages")
	repoqueryCmd.Flags().BoolP("list", "l", false, "Alias for '--files'")
	repoqueryCmd.Flags().Bool("location", false, "Like --queryformat=\"%{location}\" but deduplicated and sorted")
	repoqueryCmd.Flags().Bool("newpackage", false, "Limit to packages in newpackage advisories")
	repoqueryCmd.Flags().Bool("obsoletes", false, "Like --queryformat=\"%{obsoletes}\" but deduplicated and sorted")
	repoqueryCmd.Flags().String("providers-of", "", "After filtering is finished get selected attribute of packages and output packages that provide it")
	repoqueryCmd.Flags().Bool("provides", false, "Like --queryformat=\"%{provides}\" but deduplicated and sorted")
	repoqueryCmd.Flags().String("qf", "", "Alias for '--queryformat'")
	repoqueryCmd.Flags().String("queryformat", "", "Display format for packages. Default is \"%{full_nevra}\"")
	repoqueryCmd.Flags().Bool("querytags", false, "Display available tags for --queryformat")
	repoqueryCmd.Flags().Bool("recent", false, "Limit to only recently changed packages.")
	repoqueryCmd.Flags().Bool("recommends", false, "Like --queryformat=\"%{recommends}\" but deduplicated and sorted")
	repoqueryCmd.Flags().Bool("recursive", false, "Used with --whatrequires or --providers-of=requires options to query the packages recursively.")
	repoqueryCmd.Flags().Bool("requires", false, "Like --queryformat=\"%{requires}\" but deduplicated and sorted")
	repoqueryCmd.Flags().Bool("requires-pre", false, "Like --queryformat=\"%{requires_pre}\" but deduplicated and sorted")
	repoqueryCmd.Flags().Bool("security", false, "Limit to packages in security advisories")
	repoqueryCmd.Flags().Bool("sourcerpm", false, "Like --queryformat=\"%{sourcerpm}\" but deduplicated and sorted")
	repoqueryCmd.Flags().Bool("srpm", false, "After filtering is finished use packages' corresponding source RPMs for output")
	repoqueryCmd.Flags().Bool("suggests", false, "Like --queryformat=\"%{suggests}\" but deduplicated and sorted")
	repoqueryCmd.Flags().Bool("supplements", false, "Like --queryformat=\"%{supplements}\" but deduplicated and sorted")
	repoqueryCmd.Flags().Bool("unneeded", false, "Limit to unneeded installed packages")
	repoqueryCmd.Flags().Bool("upgrades", false, "Limit to available packages that provide an upgrade for some already installed package")
	repoqueryCmd.Flags().Bool("userinstalled", false, "Limit to packages that are not installed as dependencies or weak dependencies")
	repoqueryCmd.Flags().String("whatconflicts", "", "Limit to packages that conflict with any of <capabilities>")
	repoqueryCmd.Flags().String("whatdepends", "", "Limit to packages that require, enhance, recommend, suggest or supplement any of <capabilities>")
	repoqueryCmd.Flags().String("whatenhances", "", "Limit to packages that enhance any of <capabilities>")
	repoqueryCmd.Flags().String("whatobsoletes", "", "Limit to packages that obsolete any of <capabilities>")
	repoqueryCmd.Flags().String("whatprovides", "", "Limit to packages that provide any of <capabilities>")
	repoqueryCmd.Flags().String("whatrecommends", "", "Limit to packages that recommend any of <capabilities>")
	repoqueryCmd.Flags().String("whatrequires", "", "Limit to packages that require any of <capabilities>")
	repoqueryCmd.Flags().String("whatsuggests", "", "Limit to packages that suggest any of <capabilities>")
	repoqueryCmd.Flags().String("whatsupplements", "", "Limit to packages that supplement any of <capabilities>")
	repoqueryCmd.Flag("advisory").Hidden = true
	repoqueryCmd.Flag("bz").Hidden = true
	repoqueryCmd.Flag("cve").Hidden = true
	repoqueryCmd.Flag("list").Hidden = true
	repoqueryCmd.Flag("qf").Hidden = true
	rootCmd.AddCommand(repoqueryCmd)
}
