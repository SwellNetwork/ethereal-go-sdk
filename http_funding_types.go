package ethereal

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
