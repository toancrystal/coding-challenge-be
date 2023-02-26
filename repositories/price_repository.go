package repositories

import (
	"database/sql"
	logger "github.com/sirupsen/logrus"
	"time"
)

//go:generate mockgen -source price_repository.go -destination price_repository_mock.go -package repositories
type IPriceRepository interface {
	GetPrice(exchange string, pair string) (*Price, error)
	CreatePrice(exchange string, pair string, value string) (int64, error)
	UpdatePrice(id int, value string) (int64, error)
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

func (p *priceRepositoryImpl) CreatePrice(exchange string, pair string, value string) (int64, error) {
	current := time.Now()
	result, err := p.dbCon.Exec("insert into prices(exchange, pair, value, created_at, updated_at) values $1, $2, $3, $4, $5",
		exchange, pair, value, current, current)
	if err != nil {
		logger.Error("[PriceRepository][CreatePrice] Error while save price", err)
		return 0, err
	}

	return result.LastInsertId()
}

func (p *priceRepositoryImpl) UpdatePrice(id int, value string) (int64, error) {
	rs, err := p.dbCon.Exec("update prices set updated_at = current_timestamp, value = $1 where id = $2", value, id)
	if err != nil {
		logger.Error("[PriceRepository][CreatePrice] Error while UpdatePrice", err)
		return 0, err
	}

	return rs.RowsAffected()
}

func (p *priceRepositoryImpl) GetPrice(exchange string, pair string) (*Price, error) {
	var rs Price
	row := p.dbCon.QueryRow("select id, pair, value, exchange, updated_at from prices where exchange = $1 and pair = $2", exchange, pair)
	err := row.Scan(&rs.Id, &rs.Pair, &rs.Value, &rs.Exchange, &rs.UpdatedAt)
	if err != nil {
		logger.Errorf("[PriceRepository][GetPrice] Error while get price by exchange %s and pair %s, err=%+v",
			exchange, pair, err)
		return nil, err
	}

	return &rs, nil
}
