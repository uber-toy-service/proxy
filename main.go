package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var (
	logger     log.Logger
	db_handler *sql.DB
)

func init() {
	// initialize the logger
	var err error
	if _, err := os.Stat("logs.txt"); os.IsNotExist(err) {
		os.Create("logs.txt")
	}
	file, err := os.Open("logs.txt")
	if err != nil {
		log.Fatal("Could not open logs.txt")
	}
	var logger = log.New(file, "PROXY_LOG", log.LstdFlags)
	fmt.Fprintf(os.Stdout, "%+v\n", logger)

	// initialize the database
	db_handler, err = sql.Open("sqlite3", "meta.db")
	if err != nil {
		logger.Fatal("Not able to create Databse")
	}
}

/*
 * POST method
 * http://ServiceNameForProxy/Api[/Api]*?fwd=FwdToKwd[&ApiParam]*
 */

func ReceiveRideRequest(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {

	}
}

func main() {
	http.HandleFunc("/user", ReceiveRideRequest)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ListenAndServe did not work")
	}
}
