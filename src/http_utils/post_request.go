package http_utils

import (
	json2 "encoding/json"
	"file_utils"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time_utils"
)

// main function to handle POST requests
func Post_request_handler(w http.ResponseWriter, r *http.Request) {
	var err error
	//
	server_message := handle_post_request_data(r)
	//
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err = fmt.Fprintf(w, "%s\n", "{\"data\":\""+server_message+"\"}")
	if err != nil {
		file_utils.Log_error_to_file(err)
	}
}

// write data to file for topic
func handle_post_request_data(r *http.Request) string {
	args := r.URL.Query()
	var bodyBytes []byte
	var err error
	// read body
	bodyBytes, err = read_request_body(r)
	if err != nil {
		return "Error reading body"
	}
	// check if body has data:
	if len(bodyBytes) == 0 {
		return "No Body data"
	}
	// process body data:
	json_data, err := process_body_data(bodyBytes)
	if err != nil {
		file_utils.Log_error_to_file(err, "Error processing body data")
		return "Error processing body data"
	}
	// write json data to file:
	topic := args.Get("topic")
	if write_data_to_topic_directory(json_data, topic) {
		return "received"
	} else {
		return "write error"
	}
}

// function to read request body data
func read_request_body(r *http.Request) ([]byte, error) {
	var bodyBytes []byte
	var err error
	// read body
	if r.Body != nil {
		bodyBytes, err = ioutil.ReadAll(r.Body)
		if err != nil {
			file_utils.Log_error_to_file(err)
			return bodyBytes, err
		}

		defer func(Body io.ReadCloser) { // close body
			err = Body.Close()
			if err != nil {
				file_utils.Log_error_to_file(err)
				return
			}
		}(r.Body)
	}
	return bodyBytes, err
}

// function to process body data into json
func process_body_data(bodyBytes []byte) ([]byte, error) {
	var inter interface{}                     // interface to hold json data
	err := json2.Unmarshal(bodyBytes, &inter) // convert json to pointer
	if err != nil {
		file_utils.Log_error_to_file(err, "Error unmarshalling json")
		return nil, err
	}
	data := inter.(map[string]interface{})                          // convert pointer to map
	data["timestamp"] = time_utils.Return_current_epoch_timestamp() // add timestamp to data
	data["time_date"] = time_utils.Return_time_date_from_epoch_timestamp(data["timestamp"].(int64))
	json_data, _ := json2.Marshal(data) // convert map to json
	return json_data, err
}

// function to write data to file for topic
func write_data_to_topic_directory(json_data []byte, topic string) bool {
	cwd, _ := os.Getwd()
	path := cwd + "/topics/" + topic
	filename := file_utils.Build_file_name(topic)
	status := file_utils.Write_string_to_file(string(json_data), path, filename)
	return status
}
