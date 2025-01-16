package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/biome_completer/cmd/common"
	"github.com/spf13/cobra"
)

var formatCmd = &cobra.Command{
	Use:   "format",
	Short: "Run the formatter on a set of files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(formatCmd).Standalone()

	formatCmd.Flags().Bool("fix", false, "Alias of `--write`, writes formatted files to file system")
	formatCmd.Flags().Bool("write", false, "Writes formatted files to file system")
	common.AddGlobalFlags(formatCmd)
	common.AddVcsFlags(formatCmd)
	common.AddVcsFileFlags(formatCmd)
	common.AddFilesystemFlags(formatCmd)
	common.AddStdinFilePathFlag(formatCmd)
	common.AddFormatFlags(formatCmd)
	rootCmd.AddCommand(formatCmd)

	carapace.Gen(formatCmd).FlagCompletion(carapace.ActionMap{
		"arrow-parentheses":                 carapace.ActionValues("always", "as-needed"),
		"attribute-position":                carapace.ActionValues("multiline", "auto"),
		"bracket-same-line":                 carapace.ActionValues("true", "false"),
		"bracket-spacing":                   carapace.ActionValues("true", "false"),
		"indent-style":                      carapace.ActionValues("tab", "space"),
		"javascript-attribute-position":     carapace.ActionValues("multiline", "auto"),
		"javascript-formatter-enabled":      carapace.ActionValues("true", "false"),
		"javascript-formatter-indent-style": carapace.ActionValues("tab", "space"),
		"javascript-formatter-line-ending":  carapace.ActionValues("lf", "crlf", "cr"),
		"json-formatter-enabled":            carapace.ActionValues("true", "false"),
		"json-formatter-indent-style":       carapace.ActionValues("tab", "space"),
		"json-formatter-line-ending":        carapace.ActionValues("lf", "crlf", "cr"),
		"json-formatter-trailing-commas":    carapace.ActionValues("none", "all"),
		"jsx-quote-style":                   carapace.ActionValues("double", "single"),
		"line-ending":                       carapace.ActionValues("lf", "crlf", "cr"),
		"quote-properties":                  carapace.ActionValues("preserve", "as-needed"),
		"quote-style":                       carapace.ActionValues("double", "single"),
		"semicolons":                        carapace.ActionValues("always", "as-needed"),
		"trailing-comma":                    carapace.ActionValues("all", "es5", "none"),
		"trailing-commas":                   carapace.ActionValues("all", "es5", "none"),
		"use-editorconfig":                  carapace.ActionValues("true", "false"),
	})
}
