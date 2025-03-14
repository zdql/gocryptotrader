package key

import (
	"strings"

	"github.com/thrasher-corp/gocryptotrader/currency"
	"github.com/thrasher-corp/gocryptotrader/exchanges/asset"
)

// ExchangePairAsset is a unique map key signature for exchange, currency pair and asset
type ExchangePairAsset struct {
	Exchange string
	Base     *currency.Item
	Quote    *currency.Item
	Asset    asset.Item
}

// PairAsset is a unique map key signature for currency pair and asset
type PairAsset struct {
	Base  *currency.Item
	Quote *currency.Item
	Asset asset.Item
}

// SubAccountCurrencyAsset is a unique map key signature for subaccount, currency code and asset
type SubAccountCurrencyAsset struct {
	SubAccount string
	Currency   *currency.Item
	Asset      asset.Item
}

// MatchesExchangeAsset checks if the key matches the exchange and asset
func (k *ExchangePairAsset) MatchesExchangeAsset(exch string, item asset.Item) bool {
	return strings.EqualFold(k.Exchange, exch) && k.Asset == item
}

// MatchesPairAsset checks if the key matches the pair and asset
func (k *ExchangePairAsset) MatchesPairAsset(pair currency.Pair, item asset.Item) bool {
	return k.Base == pair.Base.Item &&
		k.Quote == pair.Quote.Item &&
		k.Asset == item
}

// MatchesExchange checks if the exchange matches
func (k *ExchangePairAsset) MatchesExchange(exch string) bool {
	return strings.EqualFold(k.Exchange, exch)
}
