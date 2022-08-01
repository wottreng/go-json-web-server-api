# go-json-web-server-api
simple json data web server written in golang 

<img src="https://github.com/wottreng/go-json-web-server-api/blob/main/pics/gopher.jpeg" alt="go-gopher"  width="auto" height="auto" />

## TLDR
a api server for sending and receiving json data

## Purpose
* handle json data per 'topic'
  * topic argument is used to select data
  * topic is an argument to the handler function
* receive and record post requests with json data
* return json data for a topic
* see `Testing` for more information

## Usage
* to run the server: `go run main.go` or just run the binary: `./main`

## Testing
* see the test file: `test_dev_server.sh` in testing directory

## Architecture
* main.go: main file
  * contains the http server
  * contains the handler function
* file_utils.go: file utilities
  * contains functions to read and write files
* time_utils.go: time utilities
  * contains functions to get the current time and date
* http_utils.go: http utilities
  * contains functions to handle http requests

## Features
* optional api key
* per topic data handling
* concurrency safe data handling


Cheers, Mark