package main

import (
	"hex/internal/adapters/app/api"
	"hex/internal/adapters/core/arithmetic"
	GRPC "hex/internal/adapters/framework/left/grpc"
	"hex/internal/adapters/framework/right/db"
	"hex/internal/ports"
	"log"
	"os"
)

func main() {

	var core ports.ArithmeticPort
	var database ports.DBPort
	var err error
	var server ports.GRPCPort
	var app ports.APIPort

	driverName := os.Getenv("DB_DRIVER_NAME")
	dataSourceName := os.Getenv("DB_DATASOURCE_NAME")
	core = arithmetic.NewAdapter()
	database, err = db.NewAdapter(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("Cannot start up the database %v", err)
	}
	defer database.CloseConnection()
	app = api.NewAdapter(core, database)
	server = GRPC.NewGRPCAdapter(app)
	server.Run()
}
