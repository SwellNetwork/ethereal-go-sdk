package ethereal

import "context"

type Client interface {
	ListProducts(ctx context.Context, params *ListProductsParams) ([]Product, error)
	ListMarketPrices(ctx context.Context, params *ListMarketPricesParams) ([]MarketPrice, error)
	GetProjectedFunding(ctx context.Context, productID string) (*ProjectedFunding, error)
	ListFunding(ctx context.Context, params ListFundingParams) ([]Funding, error)
}
