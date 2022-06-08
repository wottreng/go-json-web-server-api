package file_utils

import (
	"fmt"
	"log"
	"os"
	"strings"
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
			log.Fatal(err)
		}
	}
}

//function to create file
func CreateFile(filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}
}

func Read_string_from_file(path string, file_name string) string {
	var string_data string
	absolute_path := path + "/" + file_name
	byte_data, err := os.ReadFile(absolute_path)
	if err != nil {
		fmt.Println("[-->] File doesnt exist", err)
		topic := strings.Split(file_name, ".")
		message := "No data for: " + topic[0]
		return message
	}
	string_data = string(byte_data)
	//fmt.Println("Contents of file:", string_data)
	return string_data
}

func Write_string_to_file(data_string string, path string, file_name string) bool {
	//
	if !Does_folder_exist(path) {
		CreateFolder(path)
	}
	//
	absolute_path := path + "/" + file_name
	//
	var file_data string
	if Does_file_exist(absolute_path) {
		//println("[-->] File already exists")
		file_data = Read_string_from_file(path, file_name)
	} else {
		CreateFile(absolute_path)
	}
	data := []byte(file_data + data_string + "\n")
	err := os.WriteFile(absolute_path, data, 0644)

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("done")
	return true
}

// function for writing data to file
func Write_data_to_file(data []byte, path string, file_name string) bool {
	absolute_path := path + "/" + file_name
	file, err := os.Create(absolute_path)

	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	_, err2 := file.Write(data)

	if err2 != nil {
		log.Fatal(err2)
	}

	//fmt.Println("done")
	return true
}

// function for writing an error log
func Log_error_to_file(err error) {
	error_string := fmt.Sprintf("[-->] Error: %v", err)
	cwd, _ := os.Getwd()
	error_log_path := cwd + "/logs"
	error_log_file := "error_log.txt"
	Write_string_to_file(error_string, error_log_path, error_log_file)
}
