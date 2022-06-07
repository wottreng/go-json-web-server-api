package time_utils

import "time"

// Return_epoch_timestamp
// function to return current unix timestamp
func Return_epoch_timestamp() int64 {
	now := time.Now()
	unix_timestamp := now.Unix()
	return unix_timestamp
}

// Return_date_time_from_epoch_timestamp
// function to read unix timestamp and return date and time
func Return_date_time_from_epoch_timestamp(unix_timestamp int64) string {
	t := time.Unix(unix_timestamp, 0)
	date_time := t.Format("2006-01-02 15:04:05")
	return date_time
}
