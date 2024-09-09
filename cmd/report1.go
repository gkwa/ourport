package cmd

import (
	"github.com/gkwa/ourport/core"
	"github.com/spf13/cobra"
)

var report1Cmd = &cobra.Command{
	Use:   "report1",
	Short: "Generate report 1",
	Long:  `Generate report 1 showing statistics about image links.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := LoggerFrom(cmd.Context())
		logger.Info("Running report1 command")
		if err := core.RunReport1(); err != nil {
			logger.Error(err, "Failed to generate report1")
		}
	},
}

func init() {
	rootCmd.AddCommand(report1Cmd)
}
