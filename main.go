package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	logger log.Logger
)

func init() {
	if _, err := os.Stat("logs.txt"); os.IsNotExist(err) {
		os.Create("logs.txt")
	}
	file, err := os.Open("logs.txt")
	if err != nil {
		log.Fatal("Could not open logs.txt")
	}
	var logger = log.New(file, "PROXY_LOG", log.LstdFlags)
	fmt.Println("done with %s", logger)
}

func ReceiveRideRequest(w http.ResponseWriter, req *http.Request) {

	// We need to generate an ID for this one
	// var id int = 0
	// var foo = log.New(out io.Writer, prefix string, flag int)

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
