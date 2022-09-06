/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"database/sql"
	"github.com/spf13/cobra"
	"log"
	"market-api/internal/api/web"
	"market-api/internal/core/service"
	"market-api/internal/db/postgres"
	postgres2 "market-api/test/postgres"
	"os"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.OpenFile("testlogfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()

		log.SetOutput(f)
		log.Println("This is a test log entry")

		var db *sql.DB
		if environment == "LOCAL" {
			_, db, _ = postgres2.SetUp()
		} else {
			p := postgres.Init()
			db = p.Start()
		}

		var repository = postgres.NewMarketDb(db)
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
