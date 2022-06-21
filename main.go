/*
json web server api
TODO: add optional api key
written by Mark Wottreng
*/

package main

import (
	"fmt"
	"http_utils"
	"log"
	"net/http"
	"system_utils"
)

//http server
func main() {
	// defaults
	host_address_and_port := "localhost:8080"
	system_utils.VERBOSE = true
	// check if production mode is requested
	if system_utils.Handle_cmd_line_args() == true { // true --> prod mode
		host_address_and_port = http_utils.Return_host_ip_address_and_port()
		system_utils.VERBOSE = false
	}
	//
	println("[INFO] starting server on address: " + host_address_and_port)
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(host_address_and_port, nil))
}

//
func rootHandler(w http.ResponseWriter, r *http.Request) {
	if system_utils.VERBOSE == true {
		fmt.Printf("Headers: %+v\n", r.Header)
	}
	//check if topic arg is present
	if http_utils.Check_if_args_are_present(w, r) == false {
		return
	}

	// method handlers
	switch r.Method {
	case "GET":
		http_utils.Get_request_handler(w, r)
	case "POST":
		http_utils.Post_request_handler(w, r)
	default:
		http_utils.Method_not_allowed_handler(w)
	}
}
