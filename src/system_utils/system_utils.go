package system_utils

import (
	"fmt"
	"os"
)

// function to check cmd line args for development or production mode
func Handle_cmd_line_args() bool {
	cmd_line_args := os.Args[1:]
	//
	if len(cmd_line_args) == 0 {
		println("[INFO] no args given, enter dev mode")
		return false
	}
	//
	if cmd_line_args[0] == "-dev" {
		println("[INFO] enter dev mode")
		return false
	}
	//
	if cmd_line_args[0] == "-prod" {
		println("[INFO] enter prod mode")
		return true
	}
	//
	println("[ERROR] invalid args given!")
	println("[DEBUG] valid args: -dev or -prod")
	fmt.Println("[DEBUG] args received: ", cmd_line_args)
	println("[INFO] enter dev mode")
	return false
}
