package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Phogheus/GoBlog/goblog_services"
)

func main() {
	fmt.Print("Hi")

	http.HandleFunc("/", goblog_services.HandleRequests)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
