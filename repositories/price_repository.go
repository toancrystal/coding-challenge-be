package repositories

import (
	"database/sql"
	logger "github.com/sirupsen/logrus"
)

//go:generate mockgen -source price_repository.go -destination price_repository_mock.go -package repositories
type IPriceRepository interface {
	GetPrice(exchange string, pair string) (*Price, error)
}

func NewPriceRepository(conString string) IPriceRepository {
	dbCon := NewDatabaseConnection(conString)
	return &priceRepositoryImpl{
		dbCon: dbCon,
	}
}

type priceRepositoryImpl struct {
	dbCon *sql.DB
}

func (p *priceRepositoryImpl) GetPrice(exchange string, pair string) (*Price, error) {
	var rs Price
	err := p.dbCon.QueryRow(`select value from prices where exchange=?1 and pair=?2`, exchange, pair).Scan(&rs)
	if err != nil {
		logger.Errorf("[PriceRepository][GetPrice] Error while get price by exchange %s and pair %s, err=%+v",
			exchange, pair, err)
		return nil, err
	}

	return &rs, nil
}
