package data_utils

import (
	"file_utils"
	"os"
	"strings"
	"system_utils"
)

// read data from file
func Get_JSON_data(topic string, rows int) string {
	cwd, _ := os.Getwd()
	directory_path := cwd + "/topics/" + topic
	// check for existence of directory
	if file_utils.Does_folder_exist(directory_path) == false {
		return "{\"data\": \"No data\"}"
	}
	// return list of files in directory
	files := system_utils.List_files_in_directory(directory_path, topic)
	if len(files) == 0 {
		return "{\"data\": \"No data\"}"
	}
	// if rows = 0 then return last line in latest file
	if rows == 0 {
		filename := system_utils.Get_latest_file_in_directory(directory_path, topic)
		data_string := file_utils.Read_string_from_file(directory_path, filename)
		data_split := strings.Split(data_string, "\n")
		return "{\"data\":[" + data_split[len(data_split)-2] + "]}"
	}
	// if rows > 0 then return last rows in latest file
	if rows > 0 {
		filename := system_utils.Get_latest_file_in_directory(directory_path, topic)
		data_string := file_utils.Read_string_from_file(directory_path, filename)
		data_split := strings.Split(data_string, "\n")
		if (len(data_split) - 1) < rows {
			rows = len(data_split) - 1
		}
		data_split = data_split[len(data_split)-1-rows : len(data_split)-1] // bottom up
		all_data := strings.Join(data_split, ",")
		json_data := "{\"data\":[" + all_data + "]}"
		return json_data
	}
	// if rows == -1 then return all lines in latest file
	if rows == -1 {
		filename := system_utils.Get_latest_file_in_directory(directory_path, topic)
		data_string := file_utils.Read_string_from_file(directory_path, filename)
		data_split := strings.Split(data_string, "\n")
		data_split = data_split[:len(data_split)-1] // remove last empty line
		all_data := strings.Join(data_split, ",")
		json_data := "{\"data\":[" + all_data + "]}"
		return json_data
	}
	//
	return "{\"data\": \"error\"}"
}
