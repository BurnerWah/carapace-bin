package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/biome_completer/cmd/common"
	"github.com/spf13/cobra"
)

var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Run various checks on a set of files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lintCmd).Standalone()

	lintCmd.Flags().Bool("apply", false, "Alias for `--write`, writes safe fixes (deprecated, use `--write`)")
	lintCmd.Flags().Bool("apply-unsafe", false, "Alias for `--write --unsafe`, writes safe and unsafe fixes (deprecated, use `--write --unsafe`)")
	lintCmd.Flags().Bool("fix", false, "Alias for `--write`, writes safe fixes")
	lintCmd.Flags().String("javascript-linter-enabled", "", "Control the linter for JavaScript (and its super languages) files")
	lintCmd.Flags().String("json-linter-enabled", "", "Control the linter for JSON (and its super languages) files")
	lintCmd.Flags().String("only", "", "Run only the given rule or group of rules")
	lintCmd.Flags().String("skip", "", "Skip the given rule or group of rules")
	lintCmd.Flags().Bool("unsafe", false, "Allow to do unsafe fixes, should be used with `--write` or `--fix`")
	lintCmd.Flags().Bool("write", false, "Writes safe fixes")
	common.AddGlobalFlags(lintCmd)
	common.AddVcsFlags(lintCmd)
	common.AddVcsFileFlags(lintCmd)
	common.AddFilesystemFlags(lintCmd)
	common.AddStdinFilePathFlag(lintCmd)
	rootCmd.AddCommand(lintCmd)

	carapace.Gen(lintCmd).FlagCompletion(carapace.ActionMap{
		"javascript-linter-enabled": carapace.ActionValues("true", "false"),
		"json-linter-enabled":       carapace.ActionValues("true", "false"),
	})
}
