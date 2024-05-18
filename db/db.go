package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func CloseDB(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		log.Printf("Error while CloseDB(): %s\n", err)
		return err
	}
	return nil
}

func ConnectDb() error {
	dbs, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Printf("Error while ConnectDB() : %s\n", err)
		return err
	}

	var databaseExists bool
	err = dbs.QueryRow("SELECT EXISTS(SELECT 1 FROM information_schema.SCHEMATA WHERE SCHEMA_NAME = 'test');").Scan(&databaseExists)
	if err != nil {
		log.Printf("Error while connect DB: %s\n", err)
		return err

	}

	if !databaseExists {
		_, err = dbs.Exec("CREATE database test")
		if err != nil {
			log.Printf("Error while create db: %s\n", err)
			return err
		}
	}

	_, err = dbs.Exec("USE test")
	if err != nil {
		log.Printf("Error while connect DB: %s\n", err)
		return err
	}

	_, err = dbs.Exec(`
    CREATE TABLE IF NOT EXISTS monitor (
        URL VARCHAR(255),
        TimeRequest VARCHAR(255),
        TimeResponse VARCHAR(255),
        CodeResponse VARCHAR(255),
        Duration VARCHAR(255)
    )
`)
	if err != nil {
		log.Printf("Error while create table: %s\n", err)
		return err
	}
	db = dbs

	return nil
}
func GetDB() *sql.DB {
	return db
}
