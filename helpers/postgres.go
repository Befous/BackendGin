package helpers

import (
	"database/sql"
	"fmt"

	"github.com/Befous/BackendGin/models"
	_ "github.com/lib/pq"
)

func PostgresConnect(pconn models.PostgresInfo) *sql.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		pconn.Host, pconn.User, pconn.Password, pconn.DBName, pconn.Port, pconn.SSL)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Printf("Error connecting to PostgresDB: %v", err)
	}
	return db
}
