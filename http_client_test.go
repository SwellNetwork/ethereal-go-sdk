package ethereal

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestHTTPClientIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(HTTPClientIntegrationTestSuite))
}

type HTTPClientIntegrationTestSuite struct {
	suite.Suite

	client *HTTPClient
}

func (ts *HTTPClientIntegrationTestSuite) SetupSuite() {
	config := DefaultMainnetHTTPClientConfig()

	ts.client = NewHTTPClient(config)
}

func (ts *HTTPClientIntegrationTestSuite) TestListProducts() {
	ctx := context.Background()

	result, err := ts.client.ListProducts(ctx, nil)

	ts.T().Log("list products", result)

	ts.NoError(err)
	ts.NotNil(result)
}

func (ts *HTTPClientIntegrationTestSuite) TestListMarketPrices() {
	ctx := context.Background()
	params := &ListMarketPricesParams{
		ProductIDs: []string{
			"bc7d5575-3711-4532-a000-312bfacfb767",
			"480014cc-536e-4fd4-958b-b2afcf8ce09f",
		},
	}

	result, err := ts.client.ListMarketPrices(ctx, params)

	ts.T().Log("list market prices", result)

	ts.NoError(err)
	ts.NotNil(result)
}

func (ts *HTTPClientIntegrationTestSuite) TestGetProjectedFunding() {
	ctx := context.Background()

	result, err := ts.client.GetProjectedFunding(ctx, "bc7d5575-3711-4532-a000-312bfacfb767")

	ts.T().Log("get projected funding", result)

	ts.NoError(err)
	ts.NotNil(result)
}

func (ts *HTTPClientIntegrationTestSuite) TestListFunding() {
	ctx := context.Background()

	params := ListFundingParams{
		ProductID: "bc7d5575-3711-4532-a000-312bfacfb767",
		Range:     RangeDay,
	}
	result, err := ts.client.ListFunding(ctx, params)

	ts.T().Log("list funding", result)

	ts.NoError(err)
	ts.NotNil(result)
}
