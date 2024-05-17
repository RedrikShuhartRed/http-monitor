package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/RedrikShuhartRed/http-monitor/db"
	"github.com/RedrikShuhartRed/http-monitor/handlers"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db.ConnectDb()
	dbs := db.GetDB()
	var wg sync.WaitGroup

	link := os.Args
	if len(link) < 2 {
		fmt.Println("Enter link example: http://jojo.ru")
		return
	}
	url := link[1]

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			handlers.Get(url, dbs)
			defer wg.Done()
		}()

	}
	wg.Wait()
	db.CloseDB(dbs)

}
