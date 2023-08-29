package API

import "fmt"

func API() {
	const API_key string = "1150069161:AAEMsDBiZ3OHF8z2QEntR1V6nwSl1z7ZRzA"

	fmt.Printf("res: %v", RestClient(API_key))
}