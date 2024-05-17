package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/RedrikShuhartRed/http-monitor/db"
	"github.com/RedrikShuhartRed/http-monitor/model"

	_ "github.com/go-sql-driver/mysql"
)

var NewInfo model.Info

//var dbs *sql.DB

// func init() {
// 	db.ConnectDb()
// 	dbs = db.GetDB()
// }

func Get(link string) {

	db.ConnectDb()
	dbs := db.GetDB()
	defer dbs.Close()
	NewInfo.URL = link
	TimeRequest := time.Now()
	NewInfo.TimeRequest = TimeRequest.Format("2006-01-02 15:04:05")
	resp, err := http.Get(NewInfo.URL)
	if err != nil {
		fmt.Println("Check the URL")

		return
	}

	defer resp.Body.Close()
	NewInfo.CodeResponse = resp.StatusCode
	TimeResponse := time.Now()
	NewInfo.TimeResponse = TimeResponse.Format("2006-01-02 15:04:05")
	NewInfo.Duration = TimeResponse.Sub(TimeRequest).String()
	fmt.Println(NewInfo)

	_, err = dbs.Exec("INSERT INTO monitor (URL, TimeRequest, TimeResponse, CodeResponse, Duration) VALUES (?, ?, ?,?,?)", NewInfo.URL, NewInfo.TimeRequest, NewInfo.TimeResponse, NewInfo.CodeResponse, NewInfo.Duration)
	if err != nil {
		panic(err)
	}

}
