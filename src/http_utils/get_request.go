package http_utils

import (
	"data_utils"
	"file_utils"
	"fmt"
	"net/http"
	"strconv"
)

// main function to handle GET requests
func Get_request_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	rows := 0
	var err error
	//
	args := r.URL.Query()
	topic := args.Get("topic")
	rows_arg := args.Get("rows")
	if args.Get("alldata") == "true" {
		rows_arg = "10000000"
	}
	if rows_arg != "" {
		rows, err = strconv.Atoi(rows_arg)
	}
	//
	json_data := data_utils.Get_JSON_data(topic, rows)
	_, err = fmt.Fprintf(w, "%s\n", json_data)
	if err != nil {
		file_utils.Log_error_to_file(err, "Get_request_handler")
	}
	return
}
