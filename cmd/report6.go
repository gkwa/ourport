package cmd

import (
	"github.com/gkwa/ourport/core"
	"github.com/spf13/cobra"
)

var groupsPerPage int

var report6Cmd = &cobra.Command{
	Use:   "report6",
	Short: "Generate report 6",
	Long:  `Generate report 6 showing groups of image links in markdown files, with a specified number of groups per page.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := LoggerFrom(cmd.Context())
		logger.Info("Running report6 command")
		if err := core.RunReport6(groupsPerPage); err != nil {
			logger.Error(err, "Failed to generate report6")
		}
	},
}

func init() {
	rootCmd.AddCommand(report6Cmd)
	report6Cmd.Flags().IntVarP(&groupsPerPage, "groups-per-page", "g", 5, "Number of groups per page")
}
