package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}

	//bs := make([]byte, 9999) // 9999 specifies how many slices
	//resp.Body.Read(bs)       // Read interfaces takes in an empty byte slices and populates it in memory while returning int (number of bytes)
	//// and error
	//fmt.Println(string(bs))

	// Same as above

	//io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
