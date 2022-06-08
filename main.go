package main

// TODO: add arg to tell server the number of lines to return
// TODO: write data to file per day

import (
	"http_utils"
	"log"
	"net/http"
	"os"
)

//http server
func main() {
	host_address := http_utils.Return_host_ip_address()
	if host_address == "" {
		log.Fatal("Could not find host ip address")
		os.Exit(1)
	}
	//
	public_address := host_address + ":8080"
	println("starting server on address: " + public_address)
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(public_address, nil))
}

//
func rootHandler(w http.ResponseWriter, r *http.Request) {

	//check if topic arg is present
	if http_utils.Check_if_topic_arg_is_present(w, r) == false {
		return
	}

	//fmt.Printf("Headers: %+v\n", r.Header)

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
