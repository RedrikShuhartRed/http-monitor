package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func CloseDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		panic(err)
	}
}

func ConnectDb() {
	dbs, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}

	var databaseExists bool
	err = dbs.QueryRow("SELECT EXISTS(SELECT 1 FROM information_schema.SCHEMATA WHERE SCHEMA_NAME = 'test');").Scan(&databaseExists)
	if err != nil {
		panic(err)

	}

	if !databaseExists {
		_, err = dbs.Exec("CREATE database test")
		if err != nil {
			panic(err)
		}
	}
	// _, err = dbs.Exec("CREATE database test")
	// if err != nil {
	// 	panic(err)
	// }

	_, err = dbs.Exec("USE test")
	if err != nil {
		panic(err)
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
		panic(err)
	}
	db = dbs

}
func GetDB() *sql.DB {
	return db
}
