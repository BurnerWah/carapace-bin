package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/dnf"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dnf",
	Short: "Package management utility for RHEL, Fedora, and CentOS",
	Long:  "github.com/rpm-software-management/dnf5",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.AddGroup(
		&cobra.Group{ID: "core", Title: ""},
		&cobra.Group{ID: "query", Title: ""},
		&cobra.Group{ID: "subcommands", Title: ""},
		&cobra.Group{ID: "commands", Title: ""},
	)

	rootCmd.Flags().Bool("assumeno", false, "automatically answer no for all questions")
	rootCmd.Flags().BoolP("assumeyes", "y", false, "automatically answer yes for all questions")
	rootCmd.Flags().Bool("best", false, "try the best available package versions in transactions")
	rootCmd.PersistentFlags().BoolP("cacheonly", "C", false, "Run entirely from system cache")
	rootCmd.Flags().String("comment", "", "add a comment to transaction")
	rootCmd.Flags().String("complete", "", "Completion helper")
	rootCmd.Flags().StringP("config", "c", "", "Configuration file location")
	rootCmd.Flags().Bool("debugsolver", false, "Dump detailed solving results into files")
	rootCmd.Flags().StringSlice("disable-plugin", nil, "Disable libdnf5 plugins by name")
	rootCmd.Flags().StringSlice("disable-repo", nil, "Disable repositories")
	rootCmd.Flags().StringSlice("disableplugin", nil, "Alias for '--disable-plugin'")
	rootCmd.Flags().StringSlice("disablerepo", nil, "Alias for '--disable-repo'")
	rootCmd.Flags().Bool("dump-main-config", false, "Print main config values")
	rootCmd.Flags().String("dump-repo-config", "", "Print repo config values")
	rootCmd.Flags().Bool("dump-variables", false, "Print variable values")
	rootCmd.Flags().StringSlice("enable-plugin", nil, "Enable libdnf5 plugins by name")
	rootCmd.Flags().StringSlice("enable-repo", nil, "Enable additional repositories")
	rootCmd.Flags().StringSlice("enableplugin", nil, "Alias for '--enable-plugin'")
	rootCmd.Flags().StringSlice("enablerepo", nil, "Alias for '--enable-repo'")
	rootCmd.Flags().StringP("exclude", "x", "", "exclude packages by name or glob")
	rootCmd.Flags().String("forcearch", "", "Force the use of a different architecture")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Print help")
	rootCmd.Flags().String("installroot", "", "set install root")
	rootCmd.Flags().Bool("no-best", false, "do not limit the transaction to the best candidate")
	rootCmd.Flags().Bool("no-docs", false, "Don't install files that are marked as documentation")
	rootCmd.Flags().Bool("no-gpgchecks", false, "disable gpg signature checking (if RPM policy allows)")
	rootCmd.Flags().Bool("no-plugins", false, "Disable all libdnf5 plugins")
	rootCmd.Flags().Bool("nobest", false, "Alias for '--no-best'")
	rootCmd.Flags().Bool("nodocs", false, "Alias for '--no-docs'")
	rootCmd.Flags().Bool("nogpgcheck", false, "Alias for '--no-gpgchecks'")
	rootCmd.Flags().Bool("noplugins", false, "Alias for '--no-plugins'")
	rootCmd.Flags().BoolP("quiet", "q", false, "Quiet mode")
	rootCmd.PersistentFlags().Bool("refresh", false, "Force refreshing metadata before running the command")
	rootCmd.Flags().String("releasever", "", "override the value of $releasever in config and repo files")
	rootCmd.Flags().StringSlice("repo", nil, "Enable just specific repositories")
	rootCmd.Flags().String("repofrompath", "", "create additional repository using id and path")
	rootCmd.Flags().StringSlice("repoid", nil, "Alias for '--repo'")
	rootCmd.Flags().String("setopt", "", "set arbitrary config and repo options")
	rootCmd.Flags().String("setvar", "", "set arbitrary variable")
	rootCmd.Flags().Bool("show-new-leaves", false, "Show newly installed leaf packages and packages that became leaves after a transaction")
	rootCmd.Flags().Bool("use-host-config", false, "use config, repos, and vars from the host system rather than the installroot")
	rootCmd.Flags().Bool("version", false, "Show DNF version and exit")
	rootCmd.Flag("complete").Hidden = true
	rootCmd.Flag("disableplugin").Hidden = true
	rootCmd.Flag("disablerepo").Hidden = true
	rootCmd.Flag("enableplugin").Hidden = true
	rootCmd.Flag("enablerepo").Hidden = true
	rootCmd.Flag("nobest").Hidden = true
	rootCmd.Flag("nodocs").Hidden = true
	rootCmd.Flag("nogpgcheck").Hidden = true
	rootCmd.Flag("noplugins").Hidden = true
	rootCmd.Flag("repoid").Hidden = true

	rootCmd.MarkFlagsMutuallyExclusive("cacheonly", "refresh")
	rootCmd.MarkFlagsMutuallyExclusive("assumeyes", "assumeno") // dnf allows this but it makes no sense
	rootCmd.MarkFlagsMutuallyExclusive("best", "no-best")
	rootCmd.MarkFlagsMutuallyExclusive("enable-plugin", "no-plugins")
	rootCmd.MarkFlagsMutuallyExclusive("disable-plugin", "no-plugins")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config":           carapace.ActionFiles(),
		"disable-repo":     dnf.ActionRepos().UniqueList(","),
		"dump-repo-config": dnf.ActionRepos().UniqueList(","),
		"enable-repo":      dnf.ActionRepos().UniqueList(","),
		"repo":             dnf.ActionRepos().UniqueList(","),
	})

}
