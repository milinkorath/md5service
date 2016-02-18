// Service used to return md5 checksum of the request body
//curl --data "helloworld" http://localhost:8081
package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {

	http.HandleFunc("/", md5Sum)
	http.ListenAndServe(":8081", nil)

}
func md5Sum(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		io.WriteString(w, "Error has occured while reading body content")
	}
	checksum := md5.Sum(body)
	fmt.Fprintf(w, "%x\n", checksum)
}
