# go-json-web-server-api
simple json data web server written in golang

## Purpose
* handle json data per 'topic'
  * topic argument is used to select data
  * topic is an argument to the handler function
* receive post requests with json data
  * `curl -X POST -H 'Content-Type: application/json' -d '{"user":"mark","temp":"68"}' "0.0.0.0:8080?topic=rand"` 
* return json data for a topic ex. rand
  * `curl 0.0.0.0:8080/?topic=rand&alldata=false`

## Usage
* to run the server: `go run main.go` or just run the binary: `./main`

## Testing
* see the test file: `test.sh` in testing directory

## Architecture
* main.go: main file
  * contains the main function
  * contains the http server
* file_utils.go: file utilities
  * contains functions to read and write files

Cheers, Mark