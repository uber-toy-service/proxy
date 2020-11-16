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

func ReceiveRideRequest(w http.ResponseWriter, req *http.Request) {
	// We need to generate an ID for this one
	// Id should be loaded from the persistent database
	// var id int = 0
	// log.Printf("this is the message %v", )

	// receive the rest of the request
	// store the (ID, Status) pair in the DB (use persistent store)
	// and pass it to TripManger
}

func receive_trip_instance(w http.ResponseWriter, req *http.Request) {
	/*
		fmt.Println("-")
		fmt.Println(req.URL.Query())
		request_id, err := strconv.ParseInt(req.URL.Query().Get("request_id"), 10, 64)
		if err != nil {
			fmt.Println("Cannot parse request_id")
			return
		}
		client_id, err := strconv.ParseInt(req.URL.Query().Get("client_id"), 10, 64)
		if err != nil {
			fmt.Println("Cannot parse client_id")
			return
		}
		trip_id, err := strconv.ParseInt(req.URL.Query().Get("trip_id"), 10, 64)
		if err != nil {
			fmt.Println("Cannot parse trip_id")
			return
		}
		driver_id, err := strconv.ParseInt(req.URL.Query().Get("driver_id"), 10, 64)
		if err != nil {
			fmt.Println("Cannot parse driver_id")
			return
		}
		timestamp_unix_time, err := strconv.ParseInt(req.URL.Query().Get("timestamp"), 10, 64)
		timestamp := time.Unix(timestamp_unix_time, 0)
		if err != nil {
			fmt.Println("Cannot parse timestamp")
			return
		}
		fwd_to := req.URL.Query().Get("fwd_to")
		if err != nil {
			fmt.Println("Error")
			return
		}
	*/
	/*
		fmt.Println(request_id)
		fmt.Println(client_id)
		fmt.Println(trip_id)
		fmt.Println(driver_id)
		fmt.Println(timestamp)
		fmt.Fprintln(os.Stdout, "- Forward to %s", fwd_to)
	*/
	// http.Get(fmt.Sprintf("%s", service_url))
}

func main() {
	http.HandleFunc("/trip", ReceiveRideRequest)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ListenAndServe did not work")
	}
}
