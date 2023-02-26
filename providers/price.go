package providers

type Price struct {
	Result    Result    `json:"result"`
	Allowance Allowance `json:"allowance"`
}

type Result struct {
	Price float64 `json:"price"`
}

type Allowance struct {
	Cost      float64 `json:"cost"`
	Remaining float64 `json:"remaining"`
	Upgrade   string  `json:"upgrade"`
}
