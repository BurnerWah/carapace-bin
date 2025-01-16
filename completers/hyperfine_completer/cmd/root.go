package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hyperfine",
	Short: "A command-line benchmarking tool",
	Long:  "https://github.com/sharkdp/hyperfine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("N", "N", false, "An alias for '--shell=none'")
	rootCmd.Flags().StringP("cleanup", "c", "", "Execute CMD after the completion of all benchmarking runs for each individual command to be benchmarked")
	rootCmd.Flags().StringP("command-name", "n", "", "Give a meaningful name to a command")
	rootCmd.Flags().StringP("conclude", "C", "", "Execute CMD after each timing run")
	rootCmd.Flags().String("export-asciidoc", "", "Export the timing summary statistics as an AsciiDoc table to the given FILE")
	rootCmd.Flags().String("export-csv", "", "Export the timing summary statistics as CSV to the given FILE")
	rootCmd.Flags().String("export-json", "", "Export the timing summary statistics and timings of individual runs as JSON to the given FILE")
	rootCmd.Flags().String("export-markdown", "", "Export the timing summary statistics as a Markdown table to the given FILE")
	rootCmd.Flags().String("export-orgmode", "", "Export the timing summary statistics as an Emacs org-mode table to the given FILE")
	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.Flags().BoolP("ignore-failure", "i", false, "Ignore non-zero exit codes of the benchmarked programs")
	rootCmd.Flags().String("input", "null", "Control where the input of the benchmark comes from")
	rootCmd.Flags().IntP("max-runs", "M", 0, "Perform at most NUM runs for each command")
	rootCmd.Flags().IntP("min-runs", "m", 10, "Perform at least NUM runs for each command (default: 10)")
	rootCmd.Flags().String("output", "null", "Control where the output of the benchmark is redirected")
	rootCmd.Flags().StringSliceP("parameter-scan", "P", []string{}, "Perform benchmark runs for each value in the range MIN..MAX")
	rootCmd.Flags().Float64P("parameter-step-size", "D", 0, "Traverse the range MIN..MAX in steps of DELTA")
	rootCmd.Flags().StringSliceP("parameter-list", "L", []string{}, "Perform benchmark runs for each value in the comma-separated list VALUES")
	rootCmd.Flags().StringP("prepare", "p", "", "Execute CMD before each timing run")
	rootCmd.Flags().String("reference", "", "The reference command for the relative comparison of results")
	rootCmd.Flags().IntP("runs", "r", 0, "Perform exactly NUM runs for each command")
	rootCmd.Flags().StringP("setup", "s", "", "Execute CMD before each set of timing runs")
	rootCmd.Flags().StringP("shell", "S", "", "Set the shell to use for executing benchmarked commands")
	rootCmd.Flags().Bool("show-output", false, "Print the stdout and stderr of the benchmark instead of suppressing it")
	rootCmd.Flags().String("sort", "auto", "Specify the sort order of the speed comparison summary and the exported tables for markup formats")
	rootCmd.Flags().String("style", "auto", "Set output style type")
	rootCmd.Flags().StringP("time-unit", "u", "", "Set the time unit to be used")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")
	rootCmd.Flags().IntP("warmup", "w", 0, "Perform NUM warmup runs before the actual benchmark")

	rootCmd.Flag("parameter-scan").Nargs = 3
	rootCmd.Flag("parameter-list").Nargs = 2

	rootCmd.MarkFlagsMutuallyExclusive("runs", "max-runs")
	rootCmd.MarkFlagsMutuallyExclusive("runs", "min-runs")
	rootCmd.MarkFlagsMutuallyExclusive("N", "shell")
	rootCmd.MarkFlagsMutuallyExclusive("show-output", "style")
	rootCmd.MarkFlagsMutuallyExclusive("output", "show-output")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"export-asciidoc": carapace.ActionFiles(".adoc"),
		"export-csv":      carapace.ActionFiles(".csv"),
		"export-json":     carapace.ActionFiles(".json", ".jsonc"),
		"export-markdown": carapace.ActionFiles(".md", ".markdown"),
		"export-orgmode":  carapace.ActionFiles(".org"),
		"input": carapace.Batch(
			carapace.ActionValuesDescribed("null", "Read from /dev/null (the default)"),
			carapace.ActionFiles(),
		).ToA(),
		"output": carapace.Batch(
			carapace.ActionValuesDescribed(
				"null", "Redirect output to /dev/null (the default)",
				"pipe", "Feed the output through a pipe before discarding it",
				"inherit", "Don't redirect the output at all (same as '--show-output')",
			),
			carapace.ActionFiles(),
		).ToA(),
		"sort": carapace.ActionValuesDescribed(
			"auto", "the speed comparison will be ordered by time and the markup tables will be ordered by command (input order)",
			"command", "order benchmarks in the way they were specified",
			"mean-time", "order benchmarks by mean runtime",
		),
		"style": carapace.ActionValuesDescribed(
			"auto", "default",
			"basic", "disable output coloring and interactive elements",
			"full", "enable all effects even if no interactive terminal was detected",
			"nocolor", "keep the interactive output without any colors",
			"color", "keep the colors without any interactive output",
			"none", "disable all the output of the tool",
		),
		"time-unit": carapace.ActionValues("microsecond", "millisecond", "second"),
	})
}
