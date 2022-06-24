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
	if file_utils.Does_folder_exist(directory_path) == false {
		return "{\"data\": \"No data\"}"
	}
	//
	if rows >= 0 {
		var data_split []string
		all_files := system_utils.List_files_in_directory(directory_path, topic)
		if len(all_files) == 0 {
			return "{\"data\": \"No data\"}"
		}
		for i := 0; i < len(all_files); i++ {
			data_string_local := file_utils.Read_string_from_file(directory_path, all_files[i])
			data_split_local := strings.Split(data_string_local, "\n")
			data_split_local = data_split_local[0 : len(data_split_local)-1] // remove blank line at end
			data_split = append(data_split_local, data_split...)
			if rows <= len(data_split) {
				break
			}
		}
		length_of_data := len(data_split)
		data_split = data_split[length_of_data-rows : length_of_data]
		all_data := strings.Join(data_split, ",")
		json_data := "{\"data\":[" + all_data + "]}"
		return json_data
	}
	//
	return "{\"data\": \"error\"}"
}

func reverse_list(input []string) []string {
	if len(input) == 0 {
		return input
	}
	return append(reverse_list(input[1:]), input[0])
}

func print_list(input []string) {
	for i := 0; i < len(input); i++ {
		println(input[i])
	}
}
