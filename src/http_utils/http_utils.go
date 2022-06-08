package http_utils

import (
	"file_utils"
	"fmt"
	"log"
	"net"
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

// function to return host ip address and port
func Return_host_ip_address_and_port() string {
	acceptable_addresses := []string{"192.168.", "10.42."}
	host_ip_address := ""
	host_port := "8080"
	//
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal("[Error] " + err.Error() + "\n")
	}
	//
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				//println(ipnet.IP.String() + "\n")
				for _, acceptable_address := range acceptable_addresses {
					if ipnet.IP.String()[0:len(acceptable_address)] == acceptable_address {
						host_ip_address = ipnet.IP.String()
						break
					}
				}
			}
		}
	}
	if host_ip_address == "" {
		log.Fatal("[ERROR] Could not find host ip address")
	}
	//
	public_address_and_port := host_ip_address + ":" + host_port
	return public_address_and_port
}
