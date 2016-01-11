// proxiesget
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	// path to folder with data
	dataPath = "/home/itsme/work/goprojects/data/proxies/"

	// file with proxies data
	dataFile = dataPath + "proxies.txt"
)

// seed urls for load data
var seedUrl = map[string]string{
	"hidemyass_1": "http://proxylist.hidemyass.com/1#listable",
	"hidemyass_2": "http://proxylist.hidemyass.com/2#listable",
	"hidemyass_3": "http://proxylist.hidemyass.com/3#listable",
	"hidemyass_4": "http://proxylist.hidemyass.com/4#listable",
	"hidemyass_5": "http://proxylist.hidemyass.com/5#listable",
}

func main() {
	logit("Proxies get started")

	logit("load " + seedUrl["hidemyass_1"])
	var loadedHtml string = loadUrl(seedUrl["hidemyass_1"])

	logit("loaded")
	var fileOut = dataFile

	d1 := []byte(loadedHtml)
	err := ioutil.WriteFile(fileOut, d1, 0644)
	check(err)

	logit("saved")

}

func logit(status string) {
	fmt.Println(status)
}

/**
* check for error
 */
func check(e error) {
	if e != nil {
		panic(e)
	}
}

/**
* test loading url function
 */
func loadUrl(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	res := string(body)

	return res

}

/**
* load list of urls
 */
func loadUrlList(urls [2]string) {
	len := 2
	for i := 0; i < len; i++ {
		res := loadUrl(urls[i])
		fmt.Printf("%s", res)

		res = ""

		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
	}
}
