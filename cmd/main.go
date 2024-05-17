package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type RequestInfo struct {
	URL           string
	TimeRequest   time.Time
	TimeResponse  time.Time
	CodeResponse  int
	ContentLength string
}

func main() {

	// db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/http-monitor")
	// if err != nil {
	// 	panic(err)
	// }
	// db.Query("CREATE TABLE monitor(
	// 	url
	// )")

	RequestInfo := RequestInfo{}
	RequestInfo.URL = os.Args[1]
	fmt.Println(len(RequestInfo.URL))
	fmt.Println(RequestInfo.URL)
	RequestInfo.TimeRequest = time.Now()
	resp, _ := http.Get(RequestInfo.URL)
	RequestInfo.ContentLength = resp.Header.Get("Content-Length")
	RequestInfo.CodeResponse = resp.StatusCode
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	RequestInfo.TimeResponse = time.Now()
	reqDuration := RequestInfo.TimeResponse.Sub(RequestInfo.TimeRequest)
	fmt.Println(RequestInfo.TimeResponse.Format("2006-01-02 15:04:05"))
	fmt.Println(RequestInfo.TimeRequest.Format("2006-01-02 15:04:05"))
	fmt.Println(reqDuration)
	fmt.Println(RequestInfo)

}
