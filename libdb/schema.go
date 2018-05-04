package libdb

import (
	"github.com/jmoiron/sqlx"
	// Allows for SQL usage
	_ "github.com/lib/pq"

	"github.com/sirupsen/logrus"
)

// CreateDBSchema allows us to initialize the database
func CreateDBSchema() {
	db, err := sqlx.Open(Type, ConnectionURL)
	if err != nil {
		logrus.WithError(err).Fatal("failed to connect to database")
	}
	defer db.Close()

	schema := `
  CREATE TABLE IF NOT EXISTS tasks (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    cron_expression VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_by VARCHAR(255) NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_by VARCHAR(255) NOT NULL,
    last_started_at_time TIMESTAMP WITH TIME ZONE,
		last_started_by_host VARCHAR(255),
    last_completed_at_time TIMESTAMP WITH TIME ZONE,
		last_completed_by_host VARCHAR(255)
  )`

	db.MustExec(schema)

	if err != nil {
		logrus.WithError(err).Errorln("failed to create database schema")
		return
	}
}
