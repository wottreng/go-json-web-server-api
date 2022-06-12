package file_utils

import (
	"fmt"
	"os"
	"strings"
	"time_utils"
)

//function to check if file exists
func Does_file_exist(file_path string) bool {
	info, err := os.Stat(file_path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

//function to check if filder exists
func Does_folder_exist(folder_path string) bool {
	info, err := os.Stat(folder_path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

//function to create folder
func CreateFolder(folderPath string) {
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		err = os.MkdirAll(folderPath, 0755)
		if err != nil {
			Log_error_to_file(err, "CreateFolder")
		}
	}
}

//function to create file
func CreateFile(filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			Log_error_to_file(err, "CreateFile")
		}
		defer file.Close()
	}
}

func Read_string_from_file(path string, file_name string) string {
	var string_data string
	absolute_path := path + "/" + file_name
	byte_data, err := os.ReadFile(absolute_path)
	if err != nil {
		Log_error_to_file(err, "Read_string_from_file")
		topic := strings.Split(file_name, ".")
		message := "No data for: " + topic[0]
		return message
	}
	string_data = string(byte_data)
	return string_data
}

func Write_string_to_file(data_string string, path string, file_name string) bool {
	if !Does_folder_exist(path) {
		CreateFolder(path)
	}
	//
	absolute_path := path + "/" + file_name
	var file_data string
	if Does_file_exist(absolute_path) {
		file_data = Read_string_from_file(path, file_name)
	} else {
		CreateFile(absolute_path)
	}
	data := []byte(file_data + data_string + "\n")
	err := os.WriteFile(absolute_path, data, 0644)
	if err != nil {
		Log_error_to_file(err, "Write_string_to_file")
	}
	//
	return true
}

// function for writing data to file
func Write_data_to_file(data []byte, path string, file_name string) bool {
	var err error
	absolute_path := path + "/" + file_name
	file, err := os.Create(absolute_path)
	if err != nil {
		Log_error_to_file(err, "Write_data_to_file")
	}
	//
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			Log_error_to_file(err, "Write_data_to_file")
		}
	}(file)
	//
	_, err = file.Write(data)
	if err != nil {
		Log_error_to_file(err, "Write_data_to_file")
	}
	//
	return true
}

// function for writing an error log
func Log_error_to_file(err error, custom_message ...string) {
	var error_string string
	if len(custom_message) > 0 {
		error_string = fmt.Sprintf("[-->] Error: %v - %v", custom_message[0], err)
	} else {
		error_string = fmt.Sprintf("[-->] Error: %v", err)
	}
	cwd, _ := os.Getwd()
	error_log_path := cwd + "/logs"
	error_log_file := "error_log.txt"
	Write_string_to_file(error_string, error_log_path, error_log_file)
}

// function to build file name with topic and current date
func Build_file_name(topic string) string {
	filename := topic + "_" + time_utils.Return_current_date() + ".txt"
	return filename
}
