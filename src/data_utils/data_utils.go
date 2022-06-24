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
		data_split = data_split[0 : len(data_split)-1]
		return "{\"data\":[" + data_split[0] + "]}"
	}
	// if rows > 0 then return `rows` of data
	if rows > 0 {
		filename := system_utils.Get_latest_file_in_directory(directory_path, topic)
		data_string := file_utils.Read_string_from_file(directory_path, filename)
		data_split := strings.Split(data_string, "\n")
		data_split = data_split[0 : len(data_split)-1]
		if len(data_split) >= rows { // if more data than rows then return rows
			data_split = data_split[0:rows]
			all_data := strings.Join(data_split, ",")
			json_data := "{\"data\":[" + all_data + "]}"
			return json_data
		}
		// if the number of lines in the file is less than the number of rows requested
		additional_rows := rows - len(data_split)
		// get aditional rows from previous files
		all_files := system_utils.List_files_in_directory(directory_path, topic)
		for i := 1; i < len(all_files); i++ {
			data_string_local := file_utils.Read_string_from_file(directory_path, all_files[i])
			data_split_local := strings.Split(data_string_local, "\n")
			data_split_local = data_split_local[0 : len(data_split_local)-1]
			if (len(data_split_local)) < additional_rows {
				additional_rows -= (len(data_split_local))
				data_split_local = data_split_local[0:len(data_split_local)]
				data_split = append(data_split_local, data_split...)
			} else {
				data_split_local = data_split_local[0:additional_rows]
				data_split = append(data_split_local, data_split...)
				break
			}
		}
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
