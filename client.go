package crypto

import (
	"net/http"
)

type Client struct {
	client             *http.Client
	coinMarketCapToken string
	nomicsToken        string

	// Services
	Nomics        *NomicsService
	CoinMarketCap *CoinMarketCapService
	CoinGecko     *CoinGeckoService
}

type ClientOptionFunc func(*Client) error

func NewClient(options ...ClientOptionFunc) (*Client, error) {
	c := &Client{}
	c.client = &http.Client{}

	// Apply any given client options.
	for _, fn := range options {
		if fn == nil {
			continue
		}
		if err := fn(c); err != nil {
			return nil, err
		}
	}

	// Datasources/Services
	c.Nomics = &NomicsService{client: c}
	c.CoinMarketCap = &CoinMarketCapService{client: c}
	c.CoinGecko = &CoinGeckoService{client: c}

	return c, nil
}

func WithNomicsToken(token string) ClientOptionFunc {
	return func(c *Client) error {
		c.nomicsToken = token
		return nil
	}
}

func WithCoinMarketCapToken(token string) ClientOptionFunc {
	return func(c *Client) error {
		c.coinMarketCapToken = token
		return nil
	}
}
