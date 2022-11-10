package db

import (
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"

	// blank import for mysql driver
	_ "github.com/go-sql-driver/mysql"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	// Create the SQL Database
	var err error
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("Error creating the database %v", err)
	}

	// Test the database Connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database %v", err)
	}

	return &Adapter{db: db}, nil
}

func (adapter Adapter) CloseConnection() {
	err := adapter.db.Close()
	if err != nil {
		log.Fatalf("Error closing the database %v", err)
	}
}

func (adapter Adapter) AddHistory(answer int32, operation string) error {
	var err error
	queryString, args, err := sq.Insert("arith_table").Columns("time", "answer", "operation").Values(time.Now(), answer, operation).ToSql()
	if err != nil {
		return err
	}
	_, err = adapter.db.Exec(queryString, args...)
	if err != nil {
		return err
	}
	return nil
}
