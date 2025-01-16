package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

func AddVcsFlags(cmd *cobra.Command) {
	cmd.Flags().String("vcs-client-kind", "", "What VCS client to use")
	cmd.Flags().String("vcs-default-branch", "", "The main branch of the project")
	cmd.Flags().String("vcs-enabled", "", "Whether to integrate with the VCS client")
	cmd.Flags().String("vcs-root", "", "Folder to look for VCS files in")
	cmd.Flags().String("vcs-use-ignore-file", "", "Whether to use the VCS ignore file")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"vcs-client-kind":     carapace.ActionValues("git"),
		"vcs-default-branch":  git.ActionLocalBranches(),
		"vcs-enabled":         carapace.ActionValues("true", "false"),
		"vcs-root":            carapace.ActionDirectories(),
		"vcs-use-ignore-file": carapace.ActionValues("true", "false"),
	})
}

func AddVcsFileFlagsCi(cmd *cobra.Command) {
	cmd.Flags().Bool("changed", false, "Only check files that have been changed from the default branch")
	cmd.Flags().String("since", "", "Check files changed since the given ref")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"since": git.ActionRefs(git.RefOption{LocalBranches: true, Tags: true, HeadCommits: true}),
	})
}

func AddVcsFileFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("staged", false, "Check staged files")
	AddVcsFileFlagsCi(cmd)
}
