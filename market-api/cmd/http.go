/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"database/sql"
	"github.com/spf13/cobra"
	"market-api/internal/api/web"
	"market-api/internal/core/service"
	"market-api/internal/db/postgres"
	postgres2 "market-api/test/postgres"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "API Market",
	Long: "API Market ",
	Run: func(cmd *cobra.Command, args []string) {
		var database *sql.DB
		if environment == "LOCAL" {
			_, db, _ := postgres2.SetUp()
			database = db
		} else {
			p := postgres.Init()
			db, err := p.Start()
			if err != nil {
				panic(err.Error())
			}
			database = db
		}

		var repository = postgres.NewMarketDb(database)
		var marketService = service.MarketService{Persistence: repository}

		server := web.NewServer()
		server.Service = &marketService
		server.Start(port)
	},
}

var environment string
var port string

func init() {
	rootCmd.AddCommand(httpCmd)
	httpCmd.PersistentFlags().StringVarP(&environment, "environment", "e", "CONTAINER", "This flag sets the environment [LOCAL, CONTAINER]")
	httpCmd.PersistentFlags().StringVarP(&port, "port", "p", ":1323", "This flag sets the port of the server")
}
