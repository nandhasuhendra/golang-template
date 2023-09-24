package helper

import (
	"fmt"
	"os"
)

func BuildUrl(path string) string {
	uri := fmt.Sprintf("%s/%s", os.Getenv("TALENTA_API_URL"), path)
	return uri
}