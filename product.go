package ethereal

import (
	"context"
	"fmt"
	"strconv"
)

const (
	PathListProducts = "/v1/product"
)

type ListProductsResult struct {
	Data    []Product `json:"data"`
	HasNext bool      `json:"has_next"`
}

func (c *HTTPClient) ListProducts(ctx context.Context, params *ListProductsParams) ([]Product, error) {
	var queryParams map[string]string
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
