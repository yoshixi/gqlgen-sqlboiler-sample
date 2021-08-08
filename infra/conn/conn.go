package conn

import (
	"database/sql"
	"fmt"
	"os"
)

var (
	//DB ...
	DB *sql.DB
)

// SetupDB ...
func SetupDB() (*sql.DB, func()) {
  fmt.Printf("CONF %s \n", os.Getenv("POSTGRES_CONF"))

	DB, err := sql.Open("postgres", os.Getenv("POSTGRES_CONF"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "connect database failed: %v", err)
		os.Exit(1)
	}

	// maxOpenConns := 30
	// DB.SetMaxOpenConns(maxOpenConns)
	// DB.SetMaxIdleConns(maxOpenConns)
	// DB.SetConnMaxLifetime(time.Duration(maxOpenConns) * time.Second)

	return DB, func() {
		DB.Close()
	}
}
