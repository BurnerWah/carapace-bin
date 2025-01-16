package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/biome_completer/cmd/common"
	"github.com/spf13/cobra"
)

var ciCmd = &cobra.Command{
	Use:   "ci",
	Short: "Command to use in CI environments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ciCmd).Standalone()

	ciCmd.Flags().String("assists-enabled", "", "Allow to enable or disable the assists.")
	ciCmd.Flags().String("css-assists-enabled", "", "Control the assists for CSS files.")
	ciCmd.Flags().String("css-formatter-enabled", "", "Control the formatter for CSS (and its super")
	ciCmd.Flags().String("css-formatter-indent-style", "", "The indent style applied to CSS (and its super")
	ciCmd.Flags().String("css-formatter-indent-width", "", "The size of the indentation applied to CSS (and its")
	ciCmd.Flags().String("css-formatter-line-ending", "", "The type of line ending applied to CSS (and its")
	ciCmd.Flags().String("css-formatter-line-width", "", "What's the max width of a line applied to CSS (and its")
	ciCmd.Flags().String("css-formatter-quote-style", "", "The type of quotes used in CSS code. Defaults")
	ciCmd.Flags().String("css-linter-enabled", "", "Control the linter for CSS files.")
	ciCmd.Flags().String("formatter-enabled", "", "Allow to enable or disable the formatter check.")
	ciCmd.Flags().String("graphql-formatter-enabled", "", "Control the formatter for GraphQL files.")
	ciCmd.Flags().String("graphql-formatter-indent-style", "", "The indent style applied to GraphQL files.")
	ciCmd.Flags().String("graphql-formatter-indent-width", "", "The size of the indentation applied to GraphQL")
	ciCmd.Flags().String("graphql-formatter-line-ending", "", "The type of line ending applied to GraphQL")
	ciCmd.Flags().String("graphql-formatter-line-width", "", "What's the max width of a line applied to GraphQL")
	ciCmd.Flags().String("graphql-formatter-quote-style", "", "The type of quotes used in GraphQL code.")
	ciCmd.Flags().String("graphql-linter-enabled", "", "Control the formatter for GraphQL files.")
	ciCmd.Flags().String("javascript-assists-enabled", "", "Control the linter for JavaScript (and its super")
	ciCmd.Flags().String("javascript-linter-enabled", "", "Control the linter for JavaScript (and its super")
	ciCmd.Flags().String("json-assists-enabled", "", "Control the linter for JSON (and its super languages)")
	ciCmd.Flags().String("json-linter-enabled", "", "Control the linter for JSON (and its super languages)")
	ciCmd.Flags().String("linter-enabled", "", "Allow to enable or disable the linter check.")
	ciCmd.Flags().String("organize-imports-enabled", "", "Allow to enable or disable the organize imports.")
	common.AddGlobalFlags(ciCmd)
	common.AddVcsFlags(ciCmd)
	common.AddVcsFileFlagsCi(ciCmd)
	common.AddFilesystemFlags(ciCmd)
	common.AddFormatFlags(ciCmd)
	rootCmd.AddCommand(ciCmd)
}
