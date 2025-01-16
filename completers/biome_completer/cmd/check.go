package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/biome_completer/cmd/common"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Runs formatter, linter and import sorting to the requested files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkCmd).Standalone()

	checkCmd.Flags().Bool("apply", false, "Alias for `--write`, writes safe fixes, formatting and import sorting")
	checkCmd.Flags().Bool("apply-unsafe", false, "Alias for `--write --unsafe`, writes safe and unsafe fixes, formatting")
	checkCmd.Flags().String("assists-enabled", "", "Allow to enable or disable the assists.")
	checkCmd.Flags().String("css-assists-enabled", "", "Control the assists for CSS files.")
	checkCmd.Flags().String("css-formatter-enabled", "", "Control the formatter for CSS (and its super")
	checkCmd.Flags().String("css-formatter-indent-style", "", "The indent style applied to CSS (and its super")
	checkCmd.Flags().String("css-formatter-indent-width", "", "The size of the indentation applied to CSS (and its")
	checkCmd.Flags().String("css-formatter-line-ending", "", "The type of line ending applied to CSS (and its")
	checkCmd.Flags().String("css-formatter-line-width", "", "What's the max width of a line applied to CSS (and its")
	checkCmd.Flags().String("css-formatter-quote-style", "", "The type of quotes used in CSS code. Defaults")
	checkCmd.Flags().String("css-linter-enabled", "", "Control the linter for CSS files.")
	checkCmd.Flags().Bool("fix", false, "Alias for `--write`, writes safe fixes, formatting and import sorting")
	checkCmd.Flags().String("formatter-enabled", "", "Allow to enable or disable the formatter check.")
	checkCmd.Flags().String("graphql-formatter-enabled", "", "Control the formatter for GraphQL files.")
	checkCmd.Flags().String("graphql-formatter-indent-style", "", "The indent style applied to GraphQL files.")
	checkCmd.Flags().String("graphql-formatter-indent-width", "", "The size of the indentation applied to GraphQL")
	checkCmd.Flags().String("graphql-formatter-line-ending", "", "The type of line ending applied to GraphQL")
	checkCmd.Flags().String("graphql-formatter-line-width", "", "What's the max width of a line applied to GraphQL")
	checkCmd.Flags().String("graphql-formatter-quote-style", "", "The type of quotes used in GraphQL code.")
	checkCmd.Flags().String("graphql-linter-enabled", "", "Control the formatter for GraphQL files.")
	checkCmd.Flags().String("indent-size", "", "The size of the indentation, 2 by default (deprecated, use")
	checkCmd.Flags().String("indent-style", "", "The indent style.")
	checkCmd.Flags().String("indent-width", "", "The size of the indentation, 2 by default")
	checkCmd.Flags().String("javascript-assists-enabled", "", "Control the linter for JavaScript (and its super")
	checkCmd.Flags().String("javascript-linter-enabled", "", "Control the linter for JavaScript (and its super")
	checkCmd.Flags().String("json-assists-enabled", "", "Control the linter for JSON (and its super languages)")
	checkCmd.Flags().String("json-linter-enabled", "", "Control the linter for JSON (and its super languages)")
	checkCmd.Flags().String("linter-enabled", "", "Allow to enable or disable the linter check.")
	checkCmd.Flags().String("organize-imports-enabled", "", "Allow to enable or disable the organize imports.")
	checkCmd.Flags().Bool("unsafe", false, "Allow to do unsafe fixes, should be used with `--write` or `--fix`")
	checkCmd.Flags().Bool("write", false, "Writes safe fixes, formatting and import sorting")
	common.AddGlobalFlags(checkCmd)
	common.AddVcsFlags(checkCmd)
	common.AddVcsFileFlags(checkCmd)
	common.AddFilesystemFlags(checkCmd)
	common.AddStdinFilePathFlag(checkCmd)
	common.AddFormatFlags(checkCmd)
	rootCmd.AddCommand(checkCmd)
}
