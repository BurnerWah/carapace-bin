package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func AddGlobalFlags(cmd *cobra.Command) {
	cmd.Flags().String("colors", "", "Set the formatting mode for markup")
	cmd.Flags().String("diagnostic-level", "", "The level of diagnostics to show")
	cmd.Flags().Bool("error-on-warnings", false, "Exit on warning diagnostics")
	cmd.Flags().String("log-kind", "pretty", "How the log should look")
	cmd.Flags().String("log-level", "none", "The level of logging")
	cmd.Flags().String("max-diagnostics", "20", "Cap the amount of diagnostics displayed")
	cmd.Flags().Bool("no-errors-on-unmatched", false, "Silence errors from not mathing any files")
	cmd.Flags().String("reporter", "", "Allows to change how diagnostics and summary are reported")
	cmd.Flags().Bool("skip-errors", false, "Skip files containing syntax errors")
	cmd.Flags().Bool("use-server", false, "Connect to a running instance of the Biome daemon server")
	cmd.Flags().Bool("verbose", false, "Print additional diagnostics, and some diagnostics show more")

	AddConfigPathFlag(cmd)

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"colors":           carapace.ActionValues("off", "force"),
		"diagnostic-level": carapace.ActionValues("info", "warn", "error"),
		"log-kind":         carapace.ActionValues("pretty", "compact", "json"),
		"log-level":        carapace.ActionValues("none", "debug", "info", "warn", "error"),
		"reporter":         carapace.ActionValues("json", "json-pretty", "github", "junit", "summary", "gitlab"),
	})
}
