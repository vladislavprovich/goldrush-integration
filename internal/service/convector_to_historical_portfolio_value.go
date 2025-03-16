package service

import "goldrush-integration/pkg/integration/client"

type HistoricalPortfolioValueConvector struct{}

func NewHistoricalPortfolioValueConvector() *HistoricalPortfolioValueConvector {
	return &HistoricalPortfolioValueConvector{}
}

func (r *HistoricalPortfolioValueConvector) ConvectorToHistoricalPortfolioValueRequest(req *HistoricalPortfolioValueRequest) *client.ReqGetHistoricalPortfolioValueOverTime {
	return &client.ReqGetHistoricalPortfolioValueOverTime{
		WalletAddress: req.WalletAddress,
		ChainName:     req.ChainName,
	}
}

func (r *HistoricalPortfolioValueConvector) ConvectorToHistoricalPortfolioValueResponse(res *client.ResGetHistoricalPortfolioValueOverTime) *HistoricalPortfolioValueResponse {
	if res == nil {
		return nil
	}

	var items []HistoricalPortfolioItem
	for _, item := range res.Items {
		var holdings []Holding
		for _, h := range item.Holdings {
			holdings = append(holdings, Holding{
				QuoteRate: h.QuoteRate,
				Timestamp: h.Timestamp,
				Close: PricePoint{
					Balance:     h.Close.Balance,
					Quote:       h.Close.Quote,
					PrettyQuote: h.Close.PrettyQuote,
				},
				High: PricePoint{
					Balance:     h.High.Balance,
					Quote:       h.High.Quote,
					PrettyQuote: h.High.PrettyQuote,
				},
				Low: PricePoint{
					Balance:     h.Low.Balance,
					Quote:       h.Low.Quote,
					PrettyQuote: h.Low.PrettyQuote,
				},
				Open: PricePoint{
					Balance:     h.Open.Balance,
					Quote:       h.Open.Quote,
					PrettyQuote: h.Open.PrettyQuote,
				},
			})
		}

		items = append(items, HistoricalPortfolioItem{
			ContractAddress:      item.ContractAddress,
			ContractDecimals:     item.ContractDecimals,
			ContractName:         item.ContractName,
			ContractTickerSymbol: item.ContractTickerSymbol,
			LogoURL:              item.LogoURL,
			Holdings:             holdings,
		})
	}

	return &HistoricalPortfolioValueResponse{
		Address:       res.Address,
		UpdatedAt:     res.UpdatedAt,
		QuoteCurrency: res.QuoteCurrency,
		ChainID:       res.ChainID,
		ChainName:     res.ChainName,
		Items:         items,
	}
}
