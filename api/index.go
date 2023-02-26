package handler

import (
	"fmt"
	"github.com/toancrystal/coding-challenge-be/providers"
	"net/http"
)

var client = providers.NewCryptoWatchClient()

func Handler(w http.ResponseWriter, r *http.Request) {
	price, err := client.GetPrice("bitfinex", "ethusd")
	var response string
	if err != nil {
		response = fmt.Sprintf("<h1>Error %s</h1>", err)
	} else {
		response = fmt.Sprintf("<h1>Eth price %f</h1>", price.Result.Price)
	}
	fmt.Fprintf(w, response)
}
