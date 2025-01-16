package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "podlet",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("absolute-host-paths", "a", "", "Convert relative host paths to absolute paths")
	rootCmd.Flags().String("after", "", "Configure ordering dependency between units")
	rootCmd.Flags().String("before", "", "Configure ordering dependency between units")
	rootCmd.Flags().String("binds-to", "", "Similar to --requires, but when the dependency stops, this unit also stops")
	rootCmd.Flags().StringP("description", "d", "", "Add a description to the unit")
	rootCmd.Flags().StringP("file", "f", "", "Generate a file instead of printing to stdout")
	rootCmd.Flags().BoolP("install", "i", false, "Add an [Install] section to the unit")
	rootCmd.Flags().StringP("name", "n", "", "Override the name of the generated file (without the extension)")
	rootCmd.Flags().Bool("overwrite", false, "Overwrite existing files when generating a file")
	rootCmd.Flags().StringP("podman-version", "p", "", "Podman version generated Quadlet files should conform to")
	rootCmd.Flags().String("required-by", "", "Similar to --wanted-by, but adds stronger parent dependencies")
	rootCmd.Flags().String("requires", "", "Similar to --wants, but adds stronger requirement dependencies")
	rootCmd.Flags().Bool("skip-services-check", false, "Skip the check for existing services of the same name")
	rootCmd.Flags().BoolP("unit-directory", "u", false, "Generate a file in the Podman unit directory instead of printing to stdout")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")
	rootCmd.Flags().String("wanted-by", "", "Add (weak) parent dependencies to the unit")
	rootCmd.Flags().String("wants", "", "Add (weak) requirement dependencies to the unit")

	rootCmd.PersistentFlags().BoolP("help", "h", false, "Print help")
}
