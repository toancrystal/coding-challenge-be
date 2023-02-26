package handler

import (
	"fmt"
	logger "github.com/sirupsen/logrus"
	"github.com/toancrystal/coding-challenge-be/providers"
	"github.com/toancrystal/coding-challenge-be/repositories"
	"net/http"
	"os"
	"time"
)

var client = providers.NewCryptoWatchClient()
var connString = os.Getenv("POSTGRES_CON_STRING")
var priRepo = repositories.NewPriceRepository(connString)

func Handler(w http.ResponseWriter, r *http.Request) {
	var response string
	priceRecord, err := priRepo.GetPrice("bitfinex", "ethusd")
	if err != nil {
		logger.Error("Get price from db error ", err)
		price, err := client.GetPrice("bitfinex", "ethusd")
		if err != nil {
			logger.Error("Get price from api error ", err)
			response = fmt.Sprintf("<h1>Error %s</h1>", err)
		} else {
			priRepo.CreatePrice("bitfinex", "ethusd", fmt.Sprintf("%f", price.Result.Price))
			response = fmt.Sprintf("<h1>Eth price %f</h1>", price.Result.Price)
		}
	} else {
		if priceRecord.UpdatedAt.Add(5 * time.Minute).Before(time.Now().UTC()) {
			price, err := client.GetPrice("bitfinex", "ethusd")
			if err != nil {
				logger.Error("Get price from api error ", err)
			} else {
				priRepo.UpdatePrice(priceRecord.Id, fmt.Sprintf("%f", price.Result.Price))
			}
		}
		response = fmt.Sprintf("<h1>Eth price %f</h1>", priceRecord.Value)
	}
	fmt.Fprintf(w, response)
}
