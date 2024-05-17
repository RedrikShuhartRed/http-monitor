package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

type Info struct {
	URL           string
	TimeRequest   string
	TimeResponse  string
	CodeResponse  string
	ContentLength string
	Headers       string
}

func main() {

	url := os.Args[1]
	//http.HandleFunc(url, mainHandle)

	fmt.Println(len(url))
	fmt.Println(url)
	startTime := time.Now()
	resp, _ := http.Get(url)
	Headers := resp.Header
	fmt.Println("Заголовки", Headers)
	contentLength := resp.Header.Get("Content-Length")
	if contentLength != "" {
		fmt.Println("Размер ответа:", contentLength)
	} else {
		fmt.Println("Размер ответа не указан")
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	endTime := time.Now()
	reqDuration := endTime.Sub(startTime)
	fmt.Println(startTime)
	fmt.Println(endTime)
	fmt.Println(reqDuration)

}
