package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "glib-compile-schemas",
	Short: "Compile all GSettings schema files into a schema cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("allow-any-name", false, "Do not enforce key name restrictions")
	rootCmd.Flags().Bool("dry-run", false, "Do not write the gschema.compiled file")
	rootCmd.Flags().BoolP("help", "h", false, "Show help options")
	rootCmd.Flags().Bool("strict", false, "Abort on any errors in schemas")
	rootCmd.Flags().String("targetdir", "", "Where to store the gschemas.compiled file")
	rootCmd.Flags().Bool("version", false, "Show program version and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"targetdir": carapace.ActionDirectories(),
	})
	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
