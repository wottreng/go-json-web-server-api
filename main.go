package main

import (
	json2 "encoding/json"
	"file_utils"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

//http server
func main() {
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

//
func rootHandler(w http.ResponseWriter, r *http.Request) {
	args := r.URL.Query()
	if !args.Has("topic") {
		fmt.Fprintf(w, "Usage: %s?topic=<topic>\n", r.URL.Path)
		return
	}
	// GET request
	if r.Method == "GET" {
		data_string := handle_get_request(&args)
		if data_string == "No data" {
			fmt.Fprintf(w, "No data\n")
			return
		}
		// split data_string into lines
		data_split := strings.Split(data_string, "\n")
		data_split = data_split[:len(data_split)-1]
		// if all data is requested, return all lines
		if args.Has("alldata") == true {
			all_data := strings.Join(data_split, ",")
			json_data := "{\"data\":[" + all_data + "]}"
			fmt.Fprintf(w, json_data)
			return
		}
		// all data is not requested, return only the last line
		fmt.Fprintf(w, "%s\n", data_split[len(data_split)-2])
		return
	}
	// POST
	if r.Method == "POST" {
		handle_post_request(&args, r)
		fmt.Fprintf(w, "data rec\n")
		return
	}
	//
	fmt.Fprintf(w, "Method not allowed\n")
	return
}

// read data from file
func handle_get_request(args *url.Values) string {
	topic := args.Get("topic")
	cwd, _ := os.Getwd()
	path := cwd + "/topics"
	filename := topic + ".txt"
	if file_utils.FileExists(path + "/" + filename) {
		data_string := file_utils.Read_string_from_file(path, filename)
		return data_string
	} else {
		return "No data"
	}
}

// write data to file for topic
func handle_post_request(args *url.Values, r *http.Request) {

	var bodyBytes []byte
	var err error

	if r.Body != nil {
		bodyBytes, err = ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Body reading error: %v", err)
			return
		}
		defer r.Body.Close()
	}

	fmt.Printf("Headers: %+v\n", r.Header)

	if len(bodyBytes) == 0 {
		fmt.Printf("Body: No Body Supplied\n")
		return
	}

	var inter interface{}                        // interface to hold json data
	json2.Unmarshal(bodyBytes, &inter)           // convert json to pointer
	data := inter.(map[string]interface{})       // convert pointer to map
	data["timestamp"] = return_epoch_timestamp() // add timestamp to data
	json_data, _ := json2.Marshal(data)          // convert map to json

	topic := args.Get("topic")
	cwd, _ := os.Getwd()
	path := cwd + "/topics"
	filename := topic + ".txt"
	file_utils.Write_string_to_file(string(json_data), path, filename)
}

func return_epoch_timestamp() int64 {
	now := time.Now()
	unix_timestamp := now.Unix()
	return unix_timestamp
}
