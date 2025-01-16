package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func AddFormatFlags(cmd *cobra.Command) {
	cmd.Flags().String("arrow-parentheses", "always", "Whether to add non-necessary parentheses to arrow functions")
	cmd.Flags().String("attribute-position", "auto", "The attribute position style in HTMLish languages")
	cmd.Flags().String("bracket-same-line", "false", "Whether to hug the closing bracket of multiline HTML/JSX tags to the end of the last line")
	cmd.Flags().String("bracket-spacing", "true", "Whether to insert spaces around brackets in object literals")
	cmd.Flags().Int("indent-size", 2, "The size of the indentation (deprecated, use `indent-width`)")
	cmd.Flags().String("indent-style", "", "The indent style")
	cmd.Flags().Int("indent-width", 2, "The size of the indentation")
	cmd.Flags().String("javascript-attribute-position", "auto", "The attribute position style in jsx elements")
	cmd.Flags().String("javascript-formatter-enabled", "", "Control the formatter for JavaScript (and its super languages) files")
	cmd.Flags().Int("javascript-formatter-indent-size", 2, "The size of the indentation applied to JavaScript (and its super languages) files")
	cmd.Flags().String("javascript-formatter-indent-style", "", "The indent style applied to JavaScript (and its super languages) files")
	cmd.Flags().Int("javascript-formatter-indent-width", 2, "The size of the indentation applied to JavaScript (and its super languages) files")
	cmd.Flags().String("javascript-formatter-line-ending", "", "The type of line ending applied to JavaScript (and its super languages) files")
	cmd.Flags().Int("javascript-formatter-line-width", 80, "What's the max width of a line applied to JavaScript (and its super languages) files")
	cmd.Flags().String("json-formatter-enabled", "", "Control the formatter for JSON (and its super languages) files")
	cmd.Flags().Int("json-formatter-indent-size", 2, "The size of the indentation applied to JSON (and its super languages) files")
	cmd.Flags().String("json-formatter-indent-style", "", "The indent style applied to JSON (and its super languages) files")
	cmd.Flags().Int("json-formatter-indent-width", 2, "The size of the indentation applied to JSON (and its super languages) files")
	cmd.Flags().String("json-formatter-line-ending", "", "The type of line ending applied to JSON (and its super languages) files")
	cmd.Flags().Int("json-formatter-line-width", 80, "What's the max width of a line applied to JSON (and its super languages) files")
	cmd.Flags().String("json-formatter-trailing-commas", "none", "Print trailing commas wherever possible in multi-line comma-separated syntactic structures")
	cmd.Flags().String("jsx-quote-style", "double", "The type of quotes used in JSX")
	cmd.Flags().String("line-ending", "", "The type of line ending")
	cmd.Flags().Int("line-width", 80, "What's the max width of a line")
	cmd.Flags().String("quote-properties", "as-needed", "When properties in objects are quoted")
	cmd.Flags().String("quote-style", "double", "The type of quotes used in JavaScript code")
	cmd.Flags().String("semicolons", "", "When to use semicolons")
	cmd.Flags().String("trailing-comma", "all", "Print trailing commas wherever possible in multi-line comma-separated syntactic structures")
	cmd.Flags().String("trailing-commas", "all", "Print trailing commas wherever possible in multi-line comma-separated syntactic structures")
	cmd.Flags().String("use-editorconfig", "", "Use any `.editorconfig` files to configure the formatter")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
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
