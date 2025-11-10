package ethereal

import "context"

type Client interface {
	ListProducts(ctx context.Context, params *ListProductsParams) ([]Product, error)
	GetProjectedFunding(ctx context.Context, productID string) (*ProjectedFunding, error)
	ListFunding(ctx context.Context, params ListFundingParams) ([]Funding, error)
}

type ListProductsParams struct {
	Ticker string  `json:"ticker,omitempty"`
	Limit  float64 `json:"limit,omitempty"`
	Cursor string  `json:"cursor,omitempty"`
}

type EngineType int

const (
	EngineTypePerp EngineType = 0
	EngineTypeSpot EngineType = 1
)

type Product struct {
	Id                     string     `json:"id"`
	Ticker                 string     `json:"ticker"`
	DisplayTicker          string     `json:"displayTicker"`
	EngineType             EngineType `json:"engineType"`
	OnchainId              float64    `json:"onchainId"`
	BlockNumber            string     `json:"blockNumber"`
	BaseTokenAddress       string     `json:"baseTokenAddress"`
	QuoteTokenAddress      string     `json:"quoteTokenAddress"`
	BaseTokenName          string     `json:"baseTokenName"`
	QuoteTokenName         string     `json:"quoteTokenName"`
	LotSize                string     `json:"lotSize"`
	TickSize               string     `json:"tickSize"`
	MakerFee               string     `json:"makerFee"`
	TakerFee               string     `json:"takerFee"`
	MaxQuantity            string     `json:"maxQuantity"`
	MinQuantity            string     `json:"minQuantity"`
	MinPrice               string     `json:"minPrice"`
	MaxPrice               string     `json:"maxPrice"`
	Volume24H              string     `json:"volume24h"`
	CreatedAt              int64      `json:"createdAt"`
	CumulativeFundingUsd   string     `json:"cumulativeFundingUsd"`
	FundingUpdatedAt       int64      `json:"fundingUpdatedAt"`
	FundingRate1H          string     `json:"fundingRate1h"`
	OpenInterest           string     `json:"openInterest"`
	MaxLeverage            float64    `json:"maxLeverage"`
	MaxOpenInterestUsd     string     `json:"maxOpenInterestUsd"`
	MaxPositionNotionalUsd string     `json:"maxPositionNotionalUsd"`
	PythFeedId             float64    `json:"pythFeedId"`
}

type ProjectedFunding struct {
	FundingRateProjected1h string `json:"fundingRateProjected1h"`
	FundingRate1h          string `json:"fundingRate1h"`
}

type ListFundingParams struct {
	ProductID string `json:"productId"`
	Range     Range  `json:"range"`
}

type Range string

const (
	RangeDay   Range = "DAY"
	RangeWeek  Range = "WEEK"
	RangeMonth Range = "MONTH"
)

type Funding struct {
	FundingRate1h string `json:"fundingRate1h"`
	CreatedAt     int64  `json:"createdAt"`
}
