package ethereal

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

const (
	PathListProducts     = "/v1/product"
	PathListMarketPrices = "/v1/product/market-price"
)

type ListProductsResult struct {
	Data    []Product `json:"data"`
	HasNext bool      `json:"has_next"`
}

type ListMarketPriceResult struct {
	Data []MarketPrice `json:"data"`
}

func (c *HTTPClient) ListProducts(ctx context.Context, params *ListProductsParams) ([]Product, error) {
	queryParams := map[string]string{}
	if params != nil {
		if len(params.Ticker) > 0 {
			queryParams["ticker"] = params.Ticker
		}

		if params.Limit > 0 {
			queryParams["limit"] = strconv.FormatFloat(params.Limit, 'f', -1, 64)
		}

		if len(params.Cursor) > 0 {
			queryParams["cursor"] = params.Cursor
		}
	}

	var result *ListProductsResult
	resp, err := c.client.R().SetContext(ctx).SetResult(&result).SetQueryParams(queryParams).Get(PathListProducts)
	if err != nil {
		return nil, fmt.Errorf("list products: %v", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("list products: %s", resp.Status())
	}

	return result.Data, nil
}

func (c *HTTPClient) ListMarketPrices(ctx context.Context, params *ListMarketPricesParams) ([]MarketPrice, error) {
	if params == nil {
		return nil, fmt.Errorf("ListMarketPrices: params cannot be nil")
	}

	if len(params.ProductIDs) == 0 {
		return nil, fmt.Errorf("ListMarketPrices: productIDs cannot be empty")
	}

	queryParams := map[string]string{
		"productIds": strings.Join(params.ProductIDs, ","),
	}

	var result ListMarketPriceResult

	resp, err := c.client.R().SetContext(ctx).SetQueryParams(queryParams).SetResult(&result).Get(PathListMarketPrices)
	if err != nil {
		return nil, fmt.Errorf("list market prices: %v", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("list market prices: %s", resp.Status())
	}

	return result.Data, nil
}
