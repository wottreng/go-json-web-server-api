/*
json api web server
TODO: request data based on date
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

func main() {
	// check for cmd line args
	system_utils.Handle_cmd_line_args()
	//
	if system_utils.Mode == "dev" {
		system_utils.VERBOSE = true
	} else {
		system_utils.VERBOSE = false
		system_utils.Host_address_and_port = http_utils.Return_host_ip_address_and_port()
	}
	//
	println("[INFO] starting server on address: http://" + system_utils.Host_address_and_port + "/")
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(system_utils.Host_address_and_port, nil))
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
	//check if api arg is present
	if http_utils.Check_if_api_arg_is_present(r) == false {
		// return api key error
		w.Write([]byte("{\"error\":\"api key incorrect\"}"))
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
