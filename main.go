package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var (
	logger     log.Logger
	db_handler *sql.DB

	// Keywords that are valid for forwarding
	valid_fwd_kwds []string
)

func init() {
	// Initialize the logger
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

	// Initialize the database
	db_handler, err = sql.Open("sqlite3", "meta.db")
	if err != nil {
		logger.Fatal("Not able to create Databse")
	}
	valid_fwd_kwds = []string{"driver", "user"}
}

/*
 * POST method
 * http://ServiceNameForProxy/Api[/Api]*?fwd=FwdToKwd[&ApiParam]*
 */

func ReceiveRequest(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		logger.Print("Done with that.")
	}
	// FIXME there should only be 1 query parameter, so simplify
	var params []string = strings.Split(req.URL.RawQuery, "&")
	var fwd_data []string = strings.Split(params[0], "=")
	if fwd_data[0] != "fwd" {
		logger.Fatal("Should set the first query to 'fwd'")
	}
	var found bool = false
	var fwd_dest = fwd_data[1]
	for _, v := range valid_fwd_kwds {
		if v == fwd_dest {
			found = true
		}
	}
	if found == false {
		logger.Printf("%v not found\n", fwd_dest)
	}
	var path []string = strings.Split(req.URL.Path, "/")
	var send_to string = ""
	for _, v := range path[2:] {
		send_to += v
		send_to += "/"
	}
	send_to = strings.TrimRight(send_to, "/")

	if fwd_dest == "driver" {
		fmt.Printf("Sending to driver %v\n", send_to)
		res, err := http.Post("http://localhost:8080/"+send_to, "application/json", req.Body)
		if err != nil {
			fmt.Printf("not null with %v", res)
			log.Fatal(err)
		}
	}
}

func main() {
	http.HandleFunc("/", ReceiveRequest)
	http.HandleFunc("/user/foo/bar", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Print("user raw endpoint")
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Println("ListenAndServe did not work")
	}
}
