package http_utils

import (
	"file_utils"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// main function to handle GET requests
func Get_request_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var err error
	//
	args := r.URL.Query()
	data_string := handle_get_request_data(&args)
	if data_string == "No data" {
		_, err = fmt.Fprintf(w, "%s\n", "{\"data\":\"empty\"}")
		if err != nil {
			file_utils.Log_error_to_file(err)
			return
		}
		return
	}
	// split data_string into lines
	data_split := strings.Split(data_string, "\n")
	data_split = data_split[:len(data_split)-1]
	// if all data is requested, return all lines
	if args.Get("alldata") == "true" {
		all_data := strings.Join(data_split, ",")
		json_data := "{\"data\":[" + all_data + "]}"
		_, err = fmt.Fprintf(w, "%s\n", json_data)
		if err != nil {
			file_utils.Log_error_to_file(err)
			return
		}
		return
	}
	// all data is not requested, return only the last line
	_, err = fmt.Fprintf(w, "%s\n", data_split[len(data_split)-2])
	if err != nil {
		file_utils.Log_error_to_file(err)
		return
	}
	return
}

// read data from file
func handle_get_request_data(args *url.Values) string {
	topic := args.Get("topic")
	cwd, _ := os.Getwd()
	path := cwd + "/topics/" + topic
	filename := build_file_name(topic)
	if file_utils.Does_file_exist(path + "/" + filename) {
		data_string := file_utils.Read_string_from_file(path, filename)
		return data_string
	}
	// file does not exist, return no data
	return "No data"
}
