package repositories

import (
	"fmt"
	"os"
	"testing"
)

func TestPriceRepositoryImpl_GetPrice(t *testing.T) {
	var connString = os.Getenv("POSTGRES_CON_STRING")
	repo := NewPriceRepository(connString)
	price, err := repo.GetPrice("bitfinex", "ethusd")
	if err != nil {
		panic(err)
	}
	fmt.Println(price)
}

func TestMockIPriceRepository_CreatePrice(t *testing.T) {
	var connString = os.Getenv("POSTGRES_CON_STRING")
	repo := NewPriceRepository(connString)
	priceId, err := repo.CreatePrice("bitfinex", "ethusd", "122.2")
	if err != nil {
		panic(err)
	}
	fmt.Println(priceId)
}

func TestPriceRepositoryImpl_UpdatePrice(t *testing.T) {
	var connString = os.Getenv("POSTGRES_CON_STRING")
	repo := NewPriceRepository(connString)
	priceId, err := repo.UpdatePrice(1, "122.2")
	if err != nil {
		panic(err)
	}
	fmt.Println(priceId)
}
