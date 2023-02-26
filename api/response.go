package handler

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
