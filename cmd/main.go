package main

import (
	GRPC "github.com/ejedavy/hexGRPC/internal/adapters/left/grpc"
	"github.com/ejedavy/hexGRPC/internal/adapters/right/db"
	"github.com/ejedavy/hexGRPC/internal/app"
	"github.com/ejedavy/hexGRPC/internal/ports"
	"log"
	"os"
)

func main() {

	var database ports.DBPort
	var err error

	driverName := os.Getenv("DB_DRIVER")
	dataSourceName := os.Getenv("DS_NAME")

	// create driven adapters
	database, err = db.NewAdapter(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("Cannot start up the database %v", err)
	}
	defer database.CloseConnection()

	// plug into application ports
	application := app.New(database)

	// use application to run driver ports

	server := GRPC.NewGRPCAdapter(application)
	server.Run()
}
