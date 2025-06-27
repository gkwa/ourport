package cmd

import (
	"context"
	"os"

	"github.com/gkwa/ourport/internal/logger"
	"github.com/gkwa/ourport/version"
	"github.com/go-logr/logr"
	"github.com/spf13/cobra"
)

var (
	verbose     bool
	jsonLogging bool
)

type loggerKeyType struct{}

var loggerKey = loggerKeyType{}

func LoggerFrom(ctx context.Context) logr.Logger {
	if logger, ok := ctx.Value(loggerKey).(logr.Logger); ok {
		return logger
	}
	return logr.Discard()
}

var rootCmd = &cobra.Command{
	Use:   "ourport",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logger := logger.NewConsoleLogger(verbose, jsonLogging)
		cmd.SetContext(context.WithValue(cmd.Context(), loggerKey, logger))
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().BoolVar(&jsonLogging, "json", false, "json log output")
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `All software has versions. This is ourport's`,
	Run: func(cmd *cobra.Command, args []string) {
		buildInfo := version.GetBuildInfo()
		cmd.Println(buildInfo.String())
	},
}
