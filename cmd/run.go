package cmd

import (
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

	viper.BindPFlag("port", runCmd.Flags().Lookup("port"))
	viper.BindPFlag("data", runCmd.Flags().Lookup("data"))
}

func run(cmd *cobra.Command, args []string) {
	db.Initialize()
	http.Serve()
}
