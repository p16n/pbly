package cmd

import (
	"log"

	"github.com/p16n/pbdb/db"
	"github.com/p16n/pbly/http"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the pbly server",
	Run:   run,
}

func init() {
	viper.SetDefault("port", "8080")
	viper.SetDefault("data", "/etc/pbly/data")

	runCmd.Flags().StringP("port", "p", viper.GetString("port"), "Port pbly will bind on")
	runCmd.Flags().StringP("data", "d", viper.GetString("data"), "Data file path for pbly to use")
	runCmd.Flags().StringP("token", "t", viper.GetString("token"), "Authorization token for adding URLs")

	viper.BindPFlag("port", runCmd.Flags().Lookup("port"))
	viper.BindPFlag("data", runCmd.Flags().Lookup("data"))
	viper.BindPFlag("token", runCmd.Flags().Lookup("token"))
}

func run(cmd *cobra.Command, args []string) {
	if !viper.IsSet("token") {
		log.Fatalln("Token is not set (--token/-t TOKEN)")
	}

	db.Initialize()
	http.Serve()
}
