package handler

import (
	"encoding/json"
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
	var response PriceResponse
	priceRecord, err := priRepo.GetPrice("bitfinex", "ethusd")
	if err != nil {
		logger.Error("Get price from db error ", err)
		price, err := client.GetPrice("bitfinex", "ethusd")
		if err != nil {
			logger.Error("Get price from api error ", err)
			response = PriceResponse{
				Meta: Meta{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			}
		} else {
			priRepo.CreatePrice("bitfinex", "ethusd", fmt.Sprintf("%f", price.Result.Price))
			response = PriceResponse{
				Meta: Meta{
					Code:    http.StatusOK,
					Message: "Success",
				},
				Data: Data{
					Price: fmt.Sprintf("%f", price.Result.Price),
				},
			}
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
		response = PriceResponse{
			Meta: Meta{
				Code:    http.StatusOK,
				Message: "Success",
			},
			Data: Data{
				Price: priceRecord.Value,
			},
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

type PriceResponse struct {
	Data Data `json:"data"`
	Meta Meta `json:"meta"`
}

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Data struct {
	Price string `json:"price"`
}
