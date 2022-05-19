package crypto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	cgBaseUrl = "https://api.coingecko.com/api/v3"
)

type CoinGeckoService struct {
	client *Client
}

// CoinsIDOutput
// Docs: https://www.coingecko.com/en/api/documentation
type CoinsIDOutput struct {
	ID                           string    `json:"id"`
	Symbol                       string    `json:"symbol"`
	Name                         string    `json:"name"`
	BlockTimeInMinutes           int       `json:"block_time_in_minutes"`
	Categories                   []string  `json:"categories"`
	CountryOrigin                string    `json:"country_origin"`
	SentimentVotesUpPercentage   float64   `json:"sentiment_votes_up_percentage"`
	SentimentVotesDownPercentage float64   `json:"sentiment_votes_down_percentage"`
	MarketCapRank                int       `json:"market_cap_rank"`
	CoingeckoRank                int       `json:"coingecko_rank"`
	CoingeckoScore               float64   `json:"coingecko_score"`
	DeveloperScore               float64   `json:"developer_score"`
	CommunityScore               float64   `json:"community_score"`
	LiquidityScore               float64   `json:"liquidity_score"`
	PublicInterestScore          float64   `json:"public_interest_score"`
	LastUpdated                  time.Time `json:"last_updated"`
}

// CoinsID accepts a CoinGecko coin ID and returns a CoinsIDOutput struct
// Docs: https://www.coingecko.com/en/api/documentation
func (s *CoinGeckoService) CoinsID(id string) (CoinsIDOutput, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/coins/%s", cgBaseUrl, id), nil)
	if err != nil {
		return CoinsIDOutput{}, err
	}
	req.Header.Set("accept", "application/json")

	resp, err := s.client.client.Do(req)
	if err != nil {
		return CoinsIDOutput{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return CoinsIDOutput{}, err
	}

	var out CoinsIDOutput
	if err := json.Unmarshal(body, &out); err != nil {
		return CoinsIDOutput{}, err
	}

	return out, nil
}
