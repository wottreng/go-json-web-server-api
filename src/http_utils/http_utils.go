package http_utils

import (
	"file_utils"
	"fmt"
	"net/http"
	"time_utils"
)

// function to check if topic arg is present
func Check_if_topic_arg_is_present(w http.ResponseWriter, r *http.Request) bool {
	args := r.URL.Query()
	if !args.Has("topic") {
		_, err := fmt.Fprintf(w, "Usage: %s?topic=<topic>\n", r.URL.Path)
		if err != nil {
			file_utils.Log_error_to_file(err)
			return false
		}
		return false
	}
	return true
}

// return response for methods not supported
func Method_not_allowed_handler(w http.ResponseWriter) {
	_, err := fmt.Fprintf(w, "Method not allowed\n")
	if err != nil {
		file_utils.Log_error_to_file(err)
		return
	}
}

func build_file_name(topic string) string {
	filename := topic + "_" + time_utils.Return_current_date() + ".txt"
	return filename
}
