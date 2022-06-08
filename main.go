package main

// TODO: add request arg handler to tell server the number of lines to return

import (
	"http_utils"
	"log"
	"net/http"
	"system_utils"
)

//http server
func main() {
	host_address_and_port := "localhost:8080"
	if system_utils.Handle_cmd_line_args() == true { // true --> prod mode
		host_address_and_port = http_utils.Return_host_ip_address_and_port()
	}
	//
	println("[INFO] starting server on address: " + host_address_and_port)
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(host_address_and_port, nil))
}

//
func rootHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Printf("Headers: %+v\n", r.Header)

	//check if topic arg is present
	if http_utils.Check_if_topic_arg_is_present(w, r) == false {
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
