package cmd

import (
	"github.com/spf13/cobra"
	"github.com/jurevic/facegrinder/pkg/server"
	_ "github.com/spf13/viper"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start http server with configured api",
	Long:  `Starts a http server and serves the configured api`,
	Run: func(cmd *cobra.Command, args []string) {
		server.Run()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.
	//viper.SetDefault("port", "localhost:3000")
	//viper.SetDefault("log_level", "debug")
	//
	//viper.SetDefault("auth_login_url", "http://localhost:3000/login")
	//viper.SetDefault("auth_login_token_length", 8)
	//viper.SetDefault("auth_login_token_expiry", "11m")
	//viper.SetDefault("auth_jwt_secret", "random")
	//viper.SetDefault("auth_jwt_expiry", "15m")
	//viper.SetDefault("auth_jwt_refresh_expiry", "1h")
}