package cmd

import (
	"github.com/gkwa/ourport/core"
	"github.com/spf13/cobra"
)

var report5Cmd = &cobra.Command{
	Use:   "report5",
	Short: "Generate report 5",
	Long:  `Generate report 5 showing all groups ordered by the count of images in each group, from lowest to highest. Within each group, files are sorted numerically.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := LoggerFrom(cmd.Context())
		logger.Info("Running report5 command")
		if err := core.RunReport5(); err != nil {
			logger.Error(err, "Failed to generate report5")
		}
	},
}

func init() {
	rootCmd.AddCommand(report5Cmd)
}
