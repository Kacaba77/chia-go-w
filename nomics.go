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
	nomicsBaseURL   = "https://api.nomics.com/v1"
	nomicsAuthParam = "key"
)

type NomicsService struct {
	client *Client
}

// GetCurrenciesTickerOutput
// Docs: https://nomics.com/docs/#operation/getCurrenciesTicker
type GetCurrenciesTickerOutput []struct {
	ID                 string       `json:"id"`
	Currency           string       `json:"currency"`
	Symbol             string       `json:"symbol"`
	Name               string       `json:"name"`
	LogoURL            string       `json:"logo_url"`
	Status             string       `json:"status"`
	Price              string       `json:"price"`
	PriceDate          time.Time    `json:"price_date"`
	PriceTimestamp     time.Time    `json:"price_timestamp"`
	CirculatingSupply  string       `json:"circulating_supply"`
	MaxSupply          string       `json:"max_supply"`
	MarketCap          string       `json:"market_cap"`
	MarketCapDominance string       `json:"market_cap_dominance"`
	NumExchanges       string       `json:"num_exchanges"`
	NumPairs           string       `json:"num_pairs"`
	NumPairsUnmapped   string       `json:"num_pairs_unmapped"`
	FirstCandle        time.Time    `json:"first_candle"`
	FirstTrade         time.Time    `json:"first_trade"`
	FirstOrderBook     time.Time    `json:"first_order_book"`
	Rank               string       `json:"rank"`
	High               string       `json:"high"`
	HighTimestamp      time.Time    `json:"high_timestamp"`
	OneHour            intervalData `json:"1h"`
	OneDay             intervalData `json:"1d"`
	SevenDay           intervalData `json:"7d"`
	ThirtyDay          intervalData `json:"30d"`
	Three65Day         intervalData `json:"365d"`
	Ytd                intervalData `json:"ytd"`
}

type intervalData struct {
	Volume             string `json:"volume"`
	PriceChange        string `json:"price_change"`
	PriceChangePct     string `json:"price_change_pct"`
	VolumeChange       string `json:"volume_change"`
	VolumeChangePct    string `json:"volume_change_pct"`
	MarketCapChange    string `json:"market_cap_change"`
	MarketCapChangePct string `json:"market_cap_change_pct"`
}

// GetCurrenciesTicker accepts a list of tickers and returns a GetCurrenciesTickerOutput struct
func (s *NomicsService) GetCurrenciesTicker(symbols []string) (GetCurrenciesTickerOutput, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/currencies/ticker", nomicsBaseURL), nil)
	if err != nil {
		return GetCurrenciesTickerOutput{}, err
	}

	q := url.Values{}
	q.Add(nomicsAuthParam, s.client.nomicsToken)
	q.Add("ids", strings.Join(symbols, ","))
	q.Add("interval", "1h,1d,7d,30d,365d,ytd")
	req.URL.RawQuery = q.Encode()

	resp, err := s.client.client.Do(req)
	if err != nil {
		return GetCurrenciesTickerOutput{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return GetCurrenciesTickerOutput{}, err
	}

	var quotes GetCurrenciesTickerOutput
	if err := json.Unmarshal(body, &quotes); err != nil {
		return GetCurrenciesTickerOutput{}, err
	}

	return quotes, nil
}
