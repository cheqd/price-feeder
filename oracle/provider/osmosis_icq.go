package provider

import (
	"context"

	"cosmossdk.io/math"
	"github.com/ojo-network/price-feeder/oracle/types"
	"github.com/rs/zerolog"
)

var _ Provider = (*OsmosisICQProvider)(nil)

type (
	// OsmosisICQProvider defines an oracle provider implemented using
	// osmosis twap rate fetched via ICQ

	OsmosisICQProvider struct {
		logger zerolog.Logger

		priceStore
	}

	OsmosisICQTicker struct {
		Price math.LegacyDec
	}
)

func NewOsmosisICQProvider(
	ctx context.Context,
	logger zerolog.Logger,
	pairs ...types.CurrencyPair,
) (*OsmosisICQProvider, error) {
	osmosisICQLogger := logger.With().Str("provider", "osmosis-icq").Logger()

	provider := &OsmosisICQProvider{
		logger:     osmosisICQLogger,
		priceStore: newPriceStore(osmosisICQLogger),
	}
	provider.setCurrencyPairToTickerAndCandlePair(currencyPairToOsmosisICQPair)

	confirmedPairs, err := ConfirmPairAvailability(
		provider,
		ProviderOsmosisICQ,
		provider.logger,
		pairs...,
	)
	if err != nil {
		return nil, err
	}

	provider.setSubscribedPairs(confirmedPairs...)

	return provider, nil
}

func (p *OsmosisICQProvider) StartConnections() {}

// SubscribeCurrencyPairs sends the new subscription messages to the websocket
// and adds them to the providers subscribedPairs array
func (p *OsmosisICQProvider) SubscribeCurrencyPairs(cps ...types.CurrencyPair) {}

func (p *OsmosisICQProvider) SetTickerPrice(pair types.CurrencyPair, price math.LegacyDec) {
	osmosisICQPair := currencyPairToOsmosisICQPair(pair)
	p.setTickerPair(OsmosisICQTicker{price}, osmosisICQPair)
	telemetryWebsocketMessage(ProviderOsmosisICQ, MessageTypeTicker)
}

func (o OsmosisICQTicker) toTickerPrice() (types.TickerPrice, error) {
	tickerPrice := types.TickerPrice{
		Price:  o.Price,
		Volume: math.LegacyZeroDec(),
	}

	return tickerPrice, nil
}

// setSubscribedPairs sets N currency pairs to the map of subscribed pairs.
func (p *OsmosisICQProvider) setSubscribedPairs(cps ...types.CurrencyPair) {
	for _, cp := range cps {
		p.subscribedPairs[cp.String()] = cp
	}
}

// GetAvailablePairs returns all pairs to which the provider can subscribe.
// ex.: map["CHEQUSDC" => {}, "OJOUSDC" => {}].
func (p *OsmosisICQProvider) GetAvailablePairs() (map[string]struct{}, error) {
	// keep it static for now to support CHEQ/USDC pair
	availablePairs := map[string]struct{}{
		"CHEQUSDC": {},
	}

	return availablePairs, nil
}

// currencyPairToOsmosisPair receives a currency pair and return osmosis
// ticker symbol atomusdt@ticker.
func currencyPairToOsmosisICQPair(cp types.CurrencyPair) string {
	return cp.Base + "/" + cp.Quote
}
