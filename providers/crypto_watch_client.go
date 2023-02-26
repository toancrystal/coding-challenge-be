package providers

//go:generate mockgen -source crypto_watch_client.go -destination crypto_watch_client_mock.go -package providers
type ICryptoWatchClient interface {
	GetPrice(exchange string, pair string) (*Price, error)
}
