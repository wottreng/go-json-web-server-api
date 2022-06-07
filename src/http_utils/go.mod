module http_utils

replace file_utils => ../file_utils
replace time_utils => ../time_utils

go 1.18

require (
	file_utils v0.0.0-00010101000000-000000000000
	time_utils v0.0.0-00010101000000-000000000000
)