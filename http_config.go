package ethereal

import "time"

type HTTPClientConfig struct {
	BaseURL string        `json:"base_url"`
	Timeout time.Duration `json:"timeout"`
}

func DefaultMainnetHTTPClientConfig() HTTPClientConfig {
	return HTTPClientConfig{
		BaseURL: "https://api.ethereal.trade",
		Timeout: 3 * time.Second,
	}
}
