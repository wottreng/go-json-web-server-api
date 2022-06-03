package main

import (
	"bytes"
	json2 "encoding/json"
	"file_utils"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
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
	if r.Method == "GET" {
		data_string := handle_get_request(&args)
		data_split := strings.Split(data_string, "\n")
		if args.Has("alldata") {
			for _, line := range data_split {
				fmt.Fprintf(w, "%s ", line)
			}
		} else {
			fmt.Fprintf(w, "%s\n", data_split[len(data_split)-2])
		}
	} else if r.Method == "POST" {
		handle_post_request(&args, r)
		fmt.Fprintf(w, "data rec\n")
	} else {
		fmt.Fprintf(w, "Method not allowed\n")
	}
}

// read data from file
func handle_get_request(args *url.Values) string {
	topic := args.Get("topic")
	cwd, _ := os.Getwd()
	path := cwd + "/topics"
	filename := topic + ".txt"
	data_string := file_utils.Read_string_from_file(path, filename)

	return data_string
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
	var prettyJSON bytes.Buffer
	if err = json2.Compact(&prettyJSON, bodyBytes); err != nil {
		fmt.Printf("JSON parse error: %v", err)
		return
	}

	topic := args.Get("topic")
	cwd, _ := os.Getwd()
	path := cwd + "/topics"
	filename := topic + ".txt"
	file_utils.Write_string_to_file(string(prettyJSON.Bytes()), path, filename)
}
