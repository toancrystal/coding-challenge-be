package repositories

import "time"

type Price struct {
	Id        int       `json:"id"`
	Pair      string    `json:"pair"`
	Exchange  string    `json:"exchange"`
	Value     string    `json:"value"`
	UpdatedAt time.Time `json:"updated_at"`
}
