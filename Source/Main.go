package main

import (
	"github.com/AerisHQ/Applicator/Source/Commands"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	logger := log.NewWithOptions(os.Stdout, log.Options{
		Level:        log.InfoLevel,
		TimeFormat:   "2006-01-02 15:04:05",
		Prefix:       "Applicator",
		ReportCaller: true,
	})

	rootCmd := &cobra.Command{
		Use:   "applicator",
		Short: "A tool to run applications in a sandboxed environment",
		Long:  "Applicator allows you to run applications with specified permissions and in a sandboxed environment.",
	}

	rootCmd.AddCommand(Commands.Version)
	rootCmd.AddCommand(Commands.Run)

	if err := rootCmd.Execute(); err != nil {
		logger.Error("Error executing command", "error", err)
		os.Exit(1)
	}
}
