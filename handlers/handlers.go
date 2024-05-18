package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/RedrikShuhartRed/http-monitor/model"

	_ "github.com/go-sql-driver/mysql"
)

var NewInfo model.Info

func Get(link string, dbs *sql.DB) error {

	NewInfo.URL = link
	TimeRequest := time.Now()
	NewInfo.TimeRequest = TimeRequest.Format("2006-01-02 15:04:05.000")
	resp, err := http.Get(NewInfo.URL)
	if err != nil {
		log.Printf("Error while make http.Get(): %s\n", err)
		return err
	}

	defer resp.Body.Close()
	TimeResponse := time.Now()
	NewInfo.CodeResponse = resp.StatusCode

	NewInfo.TimeResponse = TimeResponse.Format("2006-01-02 15:04:05.000")
	NewInfo.Duration = TimeResponse.Sub(TimeRequest).String()

	_, err = dbs.Exec("USE test")
	if err != nil {
		log.Printf("Error while connect DB: %s\n", err)
		return err
	}
	_, err = dbs.Exec("INSERT INTO monitor (URL, TimeRequest, TimeResponse, CodeResponse, Duration) VALUES (?, ?, ?,?,?)", NewInfo.URL, NewInfo.TimeRequest, NewInfo.TimeResponse, NewInfo.CodeResponse, NewInfo.Duration)
	if err != nil {
		log.Printf("Error while insert in DB: %s\n", err)
		return err
	}
	return nil

}

func Average(url string, dbs *sql.DB) error {
	var Average float64

	_, err := dbs.Exec("USE test")
	if err != nil {
		log.Printf("Error while connect DB: %s\n", err)
		return err
	}
	avr := dbs.QueryRow("SELECT AVG (Duration) FROM monitor WHERE URL = ? ", url)
	err = avr.Scan(&Average)

	if err != nil {
		log.Printf("Error while select avg: %s\n", err)
		return err
	}
	fmt.Printf("Average response time for 10 simultaneous requests to %s: %f\n", url, Average)
	var maxTime string
	max := dbs.QueryRow("SELECT max(Duration) FROM monitor WHERE URL = ? ", url)
	err = max.Scan(&maxTime)

	if err != nil {
		log.Printf("Error while select max: %s\n", err)
		return err
	}
	fmt.Printf("Maximum responce time for 10 simultaneous requests to %s: %s\n", url, maxTime)

	var minTime string
	min := dbs.QueryRow("SELECT min(Duration) FROM monitor WHERE URL = ? ", url)
	err = min.Scan(&minTime)

	if err != nil {
		log.Printf("Error while select min: %s\n", err)
		return err
	}
	fmt.Printf("Minimum response time for 10 simultaneous requests to %s: %s\n", url, minTime)
	return nil
}
