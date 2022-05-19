# go-crypto-client

This is because Chris asked for it like he didn't think I'd actually do it. Well I sure showed him.

This is a library for interacting with crypto statistics APIs using a datasource of your choosing. There are Go API clients for many crypto APIs; this specific one is just meant to be as convenient as possible for interacting with multiple different supported APIs at one time. Currently supported datasources are:
 * CoinMarketCap
 * Nomics
 * CoinGecko

 NOTE: This library is missing many functions from each of the datasource APIs it supports. It's likely that functions will be added at an as-needed basis.

## Usage

Create a client, showing three clients just to illustrate the options for setting up different combinations of datasources, you can create a client that can call all datasources, or one, or none:

```golang
import (
    crypto "github.com/chia-network/go-crypto-client"
)

clientJustNomics, err := crypto.NewClient(WithNomicsToken("xxxxxx"))
clientJustCMC, err := crypto.NewClient(WithCoinMarketCapToken("yyyyyy"))

//Client with both Nomics and CoinMarketCap datasource
client, err := crypto.NewClient(
    crypto.WithNomicsToken("xxxxxx"), 
    crypto.WithCoinMarketCapToken("yyyyyy"),
)
```

Use the service for the datasource(s) chosen:

```golang
out, err := client.Nomics.GetCurrenciesTicker([]string{"XCH"})
out, err := client.CoinMarketCap.GetV1CryptocurrencyQuotesLatestOutput([]string{"XCH"})
out, err := client.CoinGecko.CoinsID("chia")
```

