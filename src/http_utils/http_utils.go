package http_utils

import (
	json2 "encoding/json"
	"file_utils"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time_utils"
)

// function to check if topic arg is present
func Check_if_topic_arg_is_present(w http.ResponseWriter, r *http.Request) bool {
	args := r.URL.Query()
	if !args.Has("topic") {
		_, err := fmt.Fprintf(w, "Usage: %s?topic=<topic>\n", r.URL.Path)
		if err != nil {
			file_utils.Log_error_to_file(err)
			return false
		}
		return false
	}
	return true
}

// main function to handle GET requests
func Get_request_handler(w http.ResponseWriter, r *http.Request) {
	var err error
	args := r.URL.Query()
	data_string := handle_get_request_data(&args)
	if data_string == "No data" {
		_, err = fmt.Fprintf(w, "No data\n")
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
	if args.Has("alldata") == true {
		all_data := strings.Join(data_split, ",")
		json_data := "{\"data\":[" + all_data + "]}"
		_, err = fmt.Fprintf(w, json_data)
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
	path := cwd + "/topics"
	filename := topic + ".txt"
	if file_utils.FileExists(path + "/" + filename) {
		data_string := file_utils.Read_string_from_file(path, filename)
		return data_string
	}
	// file does not exist, return no data
	return "No data"
}

// main function to handle POST requests
func Post_request_handler(w http.ResponseWriter, r *http.Request) {
	var err error
	handle_post_request_data(r)
	_, err = fmt.Fprintf(w, "data rec\n")
	if err != nil {
		file_utils.Log_error_to_file(err)
	}
}

// write data to file for topic
func handle_post_request_data(r *http.Request) {
	args := r.URL.Query()
	var bodyBytes []byte
	var err error

	if r.Body != nil {
		bodyBytes, err = ioutil.ReadAll(r.Body)
		if err != nil {
			file_utils.Log_error_to_file(err)
			return
		}

		defer func(Body io.ReadCloser) { // close body
			err = Body.Close()
			if err != nil {
				file_utils.Log_error_to_file(err)
				return
			}
		}(r.Body)
	}

	//fmt.Printf("Headers: %+v\n", r.Header)

	if len(bodyBytes) == 0 {
		fmt.Printf("Body: No Body Supplied\n")
		return
	}

	var inter interface{}                    // interface to hold json data
	err = json2.Unmarshal(bodyBytes, &inter) // convert json to pointer
	if err != nil {
		file_utils.Log_error_to_file(err)
		return
	}
	data := inter.(map[string]interface{})                  // convert pointer to map
	data["timestamp"] = time_utils.Return_epoch_timestamp() // add timestamp to data
	data["date_time"] = time_utils.Return_date_time_from_epoch_timestamp(data["timestamp"].(int64))
	json_data, _ := json2.Marshal(data) // convert map to json

	topic := args.Get("topic")
	cwd, _ := os.Getwd()
	path := cwd + "/topics"
	filename := topic + ".txt"
	file_utils.Write_string_to_file(string(json_data), path, filename)
}

// return response for methods not supported
func Method_not_allowed_handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Method not allowed\n")
	if err != nil {
		file_utils.Log_error_to_file(err)
		return
	}
}
