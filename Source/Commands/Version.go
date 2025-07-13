package Commands

import "github.com/spf13/cobra"

var Version = &cobra.Command{
	Use:     "version",
	Aliases: []string{"ver", "v"},
	Short:   "Display the version of Applicator",
	Long:    "Display the current version of Applicator, including build information and commit hash.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Applicator v0.1.0")
	},
}
