module data_utils

replace system_utils => ../system_utils
replace file_utils => ../file_utils

go 1.18

require (
	file_utils v0.0.0-00010101000000-000000000000
	system_utils v0.0.0-00010101000000-000000000000
)