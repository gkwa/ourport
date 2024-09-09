package cmd

import (
	"github.com/gkwa/ourport/core"
	"github.com/spf13/cobra"
)

var report4Cmd = &cobra.Command{
	Use:   "report4",
	Short: "Generate report 4",
	Long:  `Generate report 4 showing randomized groups and their image links.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := LoggerFrom(cmd.Context())
		logger.Info("Running report4 command")
		if err := core.RunReport4(); err != nil {
			logger.Error(err, "Failed to generate report4")
		}
	},
}

func init() {
	rootCmd.AddCommand(report4Cmd)
}
