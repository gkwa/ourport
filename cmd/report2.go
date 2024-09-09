package cmd

import (
	"github.com/gkwa/ourport/core"
	"github.com/spf13/cobra"
)

var report2Cmd = &cobra.Command{
	Use:   "report2",
	Short: "Generate report 2",
	Long:  `Generate report 2 showing groups of image links sorted alphabetically.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := LoggerFrom(cmd.Context())
		logger.Info("Running report2 command")
		if err := core.RunReport2(); err != nil {
			logger.Error(err, "Failed to generate report2")
		}
	},
}

func init() {
	rootCmd.AddCommand(report2Cmd)
}
