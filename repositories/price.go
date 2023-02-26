package repositories

type Price struct {
	Id       int    `json:"id"`
	Pair     string `json:"pair"`
	Exchange string `json:"exchange"`
	Value    string `json:"value"`
}
