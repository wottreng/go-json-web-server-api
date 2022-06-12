package system_utils

/*
purpose: functions for dealing with system level operations like file system, cmd line args
written by: Mark Wottreng
*/

import (
	"file_utils"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

var VERBOSE bool = false

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

// list all files in directory that match substring and order by mod time
func List_files_in_directory(directory string, file_name_substring string) []string {
	//
	var files []string
	//
	if file_utils.Does_folder_exist(directory) == false {
		return files
	}
	file_info, err := ioutil.ReadDir(directory)
	if err != nil {
		file_utils.Log_error_to_file(err, "List_files_in_directory")
		return files
	}
	// sort files in descending order by mod time
	sort.Slice(file_info, func(i, j int) bool {
		return file_info[i].ModTime().After(file_info[j].ModTime())
	})
	//
	for _, file := range file_info {
		if file.IsDir() {
			continue
		}
		//
		if strings.Contains(file.Name(), file_name_substring) {
			files = append(files, file.Name())
		}
	}
	//
	return files
}

// get latest file in directory and return its name
func Get_latest_file_in_directory(directory_path string, topic_name string) string {
	var latest_file string
	files := List_files_in_directory(directory_path, topic_name)
	//
	if len(files) == 0 {
		return ""
	}
	// files are listed in descending order
	latest_file = files[0]
	//
	return latest_file
}
