package crypto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	cmcBaseURL    = "https://pro-api.coinmarketcap.com"
	cmcAuthHeader = "X-CMC_PRO_API_KEY"
)

type CoinMarketCapService struct {
	client *Client
}

// GetV1CryptocurrencyQuotesLatestOutput
// Docs: https://coinmarketcap.com/api/documentation/v1/#operation/getV1CryptocurrencyQuotesLatest
type GetV1CryptocurrencyQuotesLatestOutput struct {
	Data   map[string]cmcData `json:"data"`
	Status struct {
		Timestamp    time.Time `json:"timestamp"`
		ErrorCode    int       `json:"error_code"`
		ErrorMessage string    `json:"error_message"`
		Elapsed      int       `json:"elapsed"`
		CreditCount  int       `json:"credit_count"`
	} `json:"status"`
}

// cmcData contains data from GetV1CryptocurrencyQuotesLatestOutput
type cmcData struct {
	ID                int                 `json:"id"`
	Name              string              `json:"name"`
	Symbol            string              `json:"symbol"`
	Slug              string              `json:"slug"`
	IsActive          int                 `json:"is_active"`
	IsFiat            int                 `json:"is_fiat"`
	CirculatingSupply int                 `json:"circulating_supply"`
	TotalSupply       int                 `json:"total_supply"`
	MaxSupply         int                 `json:"max_supply"`
	DateAdded         time.Time           `json:"date_added"`
	NumMarketPairs    int                 `json:"num_market_pairs"`
	CMCRank           int                 `json:"cmc_rank"`
	LastUpdated       time.Time           `json:"last_updated"`
	Tags              []string            `json:"tags"`
	Quote             map[string]cmcQuote `json:"quote"`
}

// cmcQuote details for a single quote
type cmcQuote struct {
	Price            float32   `json:"price"`
	Volume24H        float32   `json:"volume_24h"`
	PercentChange1H  float32   `json:"percent_change_1h"`
	PercentChange24H float32   `json:"percent_change_24h"`
	PercentChange7D  float32   `json:"percent_change_7d"`
	PercentChange30D float32   `json:"percent_change_30d"`
	MarketCap        float32   `json:"market_cap"`
	LastUpdated      time.Time `json:"last_updated"`
}

// GetV1CryptocurrencyQuotesLatest accepts a list of symbols and returns a GetV1CryptocurrencyQuotesLatestOutput struct
// Docs: https://coinmarketcap.com/api/documentation/v1/#operation/getV1CryptocurrencyQuotesLatest
func (s *CoinMarketCapService) GetV1CryptocurrencyQuotesLatest(symbols []string) (GetV1CryptocurrencyQuotesLatestOutput, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/cryptocurrency/quotes/latest", cmcBaseURL), nil)
	if err != nil {
		return GetV1CryptocurrencyQuotesLatestOutput{}, err
	}
	req.Header.Set("Accepts", "application/json")
	req.Header.Add(cmcAuthHeader, s.client.coinMarketCapToken)

	q := url.Values{}
	q.Add("symbol", strings.Join(symbols, ","))
	req.URL.RawQuery = q.Encode()

	resp, err := s.client.client.Do(req)
	if err != nil {
		return GetV1CryptocurrencyQuotesLatestOutput{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return GetV1CryptocurrencyQuotesLatestOutput{}, err
	}

	var quotes GetV1CryptocurrencyQuotesLatestOutput
	if err := json.Unmarshal(body, &quotes); err != nil {
		return GetV1CryptocurrencyQuotesLatestOutput{}, err
	}

	return quotes, nil
}
