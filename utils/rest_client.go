package API

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func RestClient(token string) string {

	var url string = "https://api.telegram.org/bot" + token

	res, err := http.Get(url + "/getUpdates")

	if err != nil {
		fmt.Println(err)
	} else {
		res, err := ioutil.ReadAll(res.Body)
		fmt.Printf("%v\n", err)
		fmt.Println(string(res))
	}
	return url
}

