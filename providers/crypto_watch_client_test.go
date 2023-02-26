package providers

import (
	"fmt"
	"testing"
)

func TestCryptoWatchClientImpl_GetPrice(t *testing.T) {
	client := NewCryptoWatchClient()
	price, err := client.GetPrice("bitfinex", "ethusd")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(price)
	}
}
