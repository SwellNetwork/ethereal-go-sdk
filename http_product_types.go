package ethereal

type ListProductsParams struct {
	Ticker string  `json:"ticker,omitempty"`
	Limit  float64 `json:"limit,omitempty"`
	Cursor string  `json:"cursor,omitempty"`
}

type ListMarketPricesParams struct {
	ProductIDs []string `json:"product_ids"`
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

type MarketPrice struct {
	ProductId    string `json:"productId"`
	BestAskPrice string `json:"bestAskPrice"`
	BestBidPrice string `json:"bestBidPrice"`
	OraclePrice  string `json:"oraclePrice"`
	Price24HAgo  string `json:"price24hAgo"`
}
