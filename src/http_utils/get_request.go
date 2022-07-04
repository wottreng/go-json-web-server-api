package http_utils

import (
	"data_utils"
	"file_utils"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// main function to handle GET requests
func Get_request_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//
	args := r.URL.Query()
	println("args: " + args.Encode())
	if args.Get("topic") != "" {
		return_topic_data(args, w)
	}
	if args.Get("list_topics") != "" {
		return_all_topics(w)
	}
	return
}

func return_topic_data(args url.Values, w http.ResponseWriter) {
	rows := 1
	var err error
	topic := args.Get("topic")
	rows_arg := args.Get("rows")
	if args.Get("alldata") == "true" {
		rows_arg = "1000"
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

func return_all_topics(w http.ResponseWriter) {
	topic_titles := file_utils.List_all_files_in_directory("topics")
	json_data := "{\"topics\":[" + strings.Join(topic_titles, ",") + "]}"
	_, err := fmt.Fprintf(w, "%s\n", json_data)
	if err != nil {
		file_utils.Log_error_to_file(err, "Get_request_handler")
	}
	return
}
