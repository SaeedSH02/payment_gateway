package main

import (
	"github.com/SaeedSH02/payment_gateway/config"
	logger "github.com/SaeedSH02/payment_gateway/log"
	"github.com/SaeedSH02/payment_gateway/repo"
	pg "github.com/SaeedSH02/payment_gateway/repo/postgres"
	"github.com/SaeedSH02/payment_gateway/routes"
)


func main() {
	logger.Initialize()


	pgDB, err := pg.NewPostgres(config.C.Postgres)
	if err != nil {
		panic(err)
	}

	var db repo.DB = pgDB

	// Setup routes and start the server
	r := routes.SetupRoutes(db)
	r.Run(":8080")
}