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
	fmt.Fprint(w, string(body))
	checksum := md5.Sum(body)
	fmt.Fprintf(w, "%x\n", checksum)
}
