package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Proxyload running")

	// data := loadUrl( "test")
	// fmt.Printf( "%s", data )

	var urls [2]string
	urls[0] = "http://example.com"
	urls[1] = "http://design.com"

	loadUrlList(urls)

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
