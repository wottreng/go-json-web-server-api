module main

replace file_utils => ./src/file_utils

replace time_utils => ./src/time_utils

replace http_utils => ./src/http_utils

replace system_utils => ./src/system_utils

//replace data_utils => ./src/data_utils

go 1.18

require (
	http_utils v0.0.0-00010101000000-000000000000
	system_utils v0.0.0-00010101000000-000000000000
)

require (
	file_utils v0.0.0-00010101000000-000000000000 // indirect
	time_utils v0.0.0-00010101000000-000000000000 // indirect
)
