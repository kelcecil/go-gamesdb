package thegamesdb

import(
	"strings"
	"os"
)

func CreateDirectoryIfNotExist(directory string) {
	 
}

func ConvertMapIntoGetParams(parameters map[string]string) string {
	params := make([]string, 0)
	for key, value := range parameters {
		if value != "" {
			params = append(params, key+"="+value)
		}
	}
	return strings.Join(params, "&")
}