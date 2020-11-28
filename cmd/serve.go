package cmd

import (
	"os"

	"github.com/rodopoulos/plantai/server"
	"github.com/rodopoulos/plantai/utils"
	"github.com/spf13/cobra"
)

var (
	serverHost string
	serverPort string
	herokuMode bool
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start plantai in server mode",
	Long:  "start plantai in server mode",
	Run: func(cmd *cobra.Command, args []string) {
		addr := utils.BuildServerAddress(serverHost, serverPort, herokuMode)
		logger := utils.NewLogger()
		srv := server.NewServer(addr, logger)
		if err := srv.Start(); err != nil {
			panic(err)
		}

		os.Exit(0)
	},
}

func init() {
	serveCmd.Flags().StringVarP(&serverHost, "host", "H", "localhost", "server host")
	serveCmd.Flags().StringVarP(&serverPort, "port", "p", "8080", "server port")
	serveCmd.Flags().BoolVar(&herokuMode, "heroku", false, "heroku mode")
	RootCmd.AddCommand(serveCmd)
}
