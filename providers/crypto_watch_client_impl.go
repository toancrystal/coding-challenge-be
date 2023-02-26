package providers

import (
	"encoding/json"
	"fmt"
	logger "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type cryptoWatchClientImpl struct {
}

func (c *cryptoWatchClientImpl) GetPrice(exchange string, pair string) (*Price, error) {
	requestURL := fmt.Sprintf("https://api.cryptowat.ch/markets/%s/%s/price", exchange, pair)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		logger.Error("[CryptoWatchClient][GetPrice]: could not create request, err=", err)
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Error("[CryptoWatchClient][GetPrice] get price error, err=", err)
		return nil, err
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			logger.Error("[CryptoWatchClient][GetPrice] close get price body error, err=", err)
		}
	}()

	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		logger.Error("[CryptoWatchClient][GetPrice] read get price body error: ", err)
		return nil, err
	}

	var priceResponse Price
	err = json.Unmarshal(responseData, &priceResponse)
	if err != nil {
		logger.Error("[CryptoWatchClient][GetPrice] read response body error: ", err)
		return nil, err
	}
	return &priceResponse, nil
}

func NewCryptoWatchClient() ICryptoWatchClient {
	return &cryptoWatchClientImpl{}
}
