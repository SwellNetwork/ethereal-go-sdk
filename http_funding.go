package ethereal

import (
	"context"
	"fmt"
)

const (
	PathGetProjectedFunding = "/v1/funding/projected"
	PathListFunding         = "/v1/funding"
)

type ListFundingResult struct {
	Data []Funding `json:"data"`
}

func (c *HTTPClient) GetProjectedFunding(ctx context.Context, productID string) (*ProjectedFunding, error) {
	queryParams := map[string]string{
		"productId": productID,
	}

	var result ProjectedFunding

	resp, err := c.client.R().SetContext(ctx).SetQueryParams(queryParams).SetResult(&result).Get(PathGetProjectedFunding)
	if err != nil {
		return nil, fmt.Errorf("get projected funding error: %s", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("get projected funding error: %s", resp.Status())
	}

	return &result, nil
}

func (c *HTTPClient) ListFunding(ctx context.Context, params ListFundingParams) ([]Funding, error) {
	queryParams := map[string]string{
		"productId": params.ProductID,
		"range":     string(params.Range),
	}

	var result ListFundingResult
	resp, err := c.client.R().SetContext(ctx).SetQueryParams(queryParams).SetResult(&result).Get(PathListFunding)
	if err != nil {
		return nil, fmt.Errorf("list funding error: %s", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("list funding error: %s", resp.Status())
	}

	return result.Data, nil
}
