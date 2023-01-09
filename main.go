package main

// Usage
// $ trim-amzn-url <URL>

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

func main() {
	// get arguments
	args := os.Args[1:]
	url_string := args[0]
	// get url hostname
	url, err := url.Parse(url_string)
	if err != nil {
		fmt.Println("Error parsing url")
		return
	}
	hostname := url.Hostname()
	path := url.Path
	// split path with "/"
	path_split := strings.Split(path, "/")
	// find "dp" in path
	var dp_index int
	for i, v := range path_split {
		if v == "dp" {
			dp_index = i
		}
	}

	result_path := "/dp/" + path_split[dp_index+1]
	result := "https://" + hostname + result_path

	fmt.Println(result)
}
