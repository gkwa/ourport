package cmd

import (
	"github.com/gkwa/ourport/core"
	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := LoggerFrom(cmd.Context())
		logger.Info("Running hello command")
		links, err := core.FetchImageLinks()
		if err != nil {
			logger.Error(err, "Failed to fetch image links")
			return
		}
		core.Report1(links)
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
