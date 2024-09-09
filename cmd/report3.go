package cmd

import (
	"github.com/gkwa/ourport/core"
	"github.com/spf13/cobra"
)

var report3Cmd = &cobra.Command{
	Use:   "report3",
	Short: "Generate report 3",
	Long:  `Generate report 3 showing all image links with a counter.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := LoggerFrom(cmd.Context())
		logger.Info("Running report3 command")
		if err := core.RunReport3(); err != nil {
			logger.Error(err, "Failed to generate report3")
		}
	},
}

func init() {
	rootCmd.AddCommand(report3Cmd)
}
