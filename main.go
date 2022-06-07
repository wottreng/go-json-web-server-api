package main

import (
	"http_utils"
	"log"
	"net/http"
)

//http server
func main() {
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
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
		http_utils.Method_not_allowed_handler(w, r)
	}
}
