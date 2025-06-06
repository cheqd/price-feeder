package provider

import (
	"context"
	"testing"

	"cosmossdk.io/math"
	"github.com/ojo-network/price-feeder/oracle/types"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestOsmosisICQProvider_GetTickerPrices(t *testing.T) {
	p, err := NewOsmosisICQProvider(
		context.TODO(),
		zerolog.Nop(),
		CHEQUSDC,
	)
	require.NoError(t, err)

	t.Run("valid_request_single_ticker", func(t *testing.T) {
		lastPrice := math.LegacyMustNewDecFromStr("34.69000000")
		volume := math.LegacyMustNewDecFromStr("2396974.02000000")

		tickerMap := map[string]types.TickerPrice{}
		tickerMap["CHEQ/USDC"] = types.TickerPrice{
			Price:  lastPrice,
			Volume: volume,
		}

		p.tickers = tickerMap

		prices, err := p.GetTickerPrices(CHEQUSDC)
		require.NoError(t, err)
		require.Len(t, prices, 1)
		require.Equal(t, lastPrice, prices[CHEQUSDC].Price)
		require.Equal(t, volume, prices[CHEQUSDC].Volume)
	})

	t.Run("valid_request_multi_ticker", func(t *testing.T) {
		lastPriceAtom := math.LegacyMustNewDecFromStr("34.69000000")
		lastPriceLuna := math.LegacyMustNewDecFromStr("41.35000000")
		volume := math.LegacyMustNewDecFromStr("2396974.02000000")

		tickerMap := map[string]types.TickerPrice{}
		tickerMap["CHEQ/USDC"] = types.TickerPrice{
			Price:  lastPriceAtom,
			Volume: volume,
		}

		tickerMap["LUNA/USDT"] = types.TickerPrice{
			Price:  lastPriceLuna,
			Volume: volume,
		}

		p.tickers = tickerMap
		prices, err := p.GetTickerPrices(
			CHEQUSDC,
			LUNAUSDT,
		)
		require.NoError(t, err)
		require.Len(t, prices, 2)
		require.Equal(t, lastPriceAtom, prices[CHEQUSDC].Price)
		require.Equal(t, volume, prices[CHEQUSDC].Volume)
		require.Equal(t, lastPriceLuna, prices[LUNAUSDT].Price)
		require.Equal(t, volume, prices[LUNAUSDT].Volume)
	})

	t.Run("valid_request_use_set_ticker_price", func(t *testing.T) {
		lastPrice := math.LegacyMustNewDecFromStr("34.69000000")
		p.SetTickerPrice(CHEQUSDC, lastPrice)

		prices, err := p.GetTickerPrices(CHEQUSDC)
		require.NoError(t, err)
		require.Len(t, prices, 1)
		require.Equal(t, lastPrice, prices[CHEQUSDC].Price)
		require.Equal(t, math.LegacyZeroDec(), prices[CHEQUSDC].Volume)
	})

	t.Run("invalid_request_invalid_ticker", func(t *testing.T) {
		prices, _ := p.GetTickerPrices(types.CurrencyPair{Base: "FOO", Quote: "BAR"})
		require.Empty(t, prices)
	})
}
