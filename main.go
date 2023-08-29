package main

import (
	"fmt"
	API "gotg/utils"
	"io/ioutil"
	"net/http"
)

func main() {
	API.API()

	http.HandleFunc("/bot", func (res http.ResponseWriter, req *http.Request)  {
		reqBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Printf("Error: %v", err.Error())
		} else {
			fmt.Printf("res: %v", string(reqBody))
		}
	})
	http.ListenAndServe(":8080", nil)

}