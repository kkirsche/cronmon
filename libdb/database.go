package libdb

import "fmt"

const (
	// Type is the database type used in sql.open()
	Type = "postgres"
	// Scheme is the scheme to use in the connection URL
	Scheme = "postgresql"
	// Host is the host to connect to the database on
	Host = "localhost"
	// Port is the database port to connect to
	Port = 26257
	// Username is the database user that we connect as
	Username = "evilroach"
	// Database is the actual database which we connect to
	Database = "cronmon"
	// SSLmode indicates whether we should use SSL for our database connection
	SSLmode = "disable"
)

// ConnectionURL is the actual URL which can be used to connect to the database
var ConnectionURL = fmt.Sprintf("%s://%s@%s:%d/%s?sslmode=%s",
	Scheme, Username, Host, Port, Database, SSLmode)
