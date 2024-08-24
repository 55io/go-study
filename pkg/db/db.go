package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/55io/go-study/pkg/config"
)

func New(dbConf *config.Db) *sql.DB {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbConf.Host, dbConf.Port, dbConf.User, dbConf.Password, dbConf.Dbname, dbConf.SllMode,
	)

	db, err := sql.Open(
		"postgres",
		connStr,
	)

	if err != nil {
		log.Fatal("Postgres connection error", err)
	}

	log.Println("DB connected!")

	return db
}
