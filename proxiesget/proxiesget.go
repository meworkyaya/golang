// proxiesget
package main

import (
	"fmt"
)

const (
	// path to folder with data
	dataPath = "/home/itsme/work/goprojects/data/proxies/"

	// file with proxies data
	dataFile = dataPath + "proxies.txt"
)

// seed urls for load data
var seedUrl = map[string]string{}

func main() {
	fmt.Println("Proxies get started")
}
