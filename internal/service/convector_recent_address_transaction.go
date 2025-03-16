package service

import "goldrush-integration/pkg/integration/client"

type RecentAddressTransactionConvector struct{}

func NewRecentAddressTransactionConvector() *RecentAddressTransactionConvector {
	return &RecentAddressTransactionConvector{}
}

func (r *RecentAddressTransactionConvector) ConvectorToRecentAddressTransactionRequest(req *RecentAddressTransactionRequest) *client.ReqGetRecentTransactionForAddress {
	return &client.ReqGetRecentTransactionForAddress{
		WalletAddress: req.WalletAddress,
		ChainName:     req.ChainName,
	}
}

func (r *RecentAddressTransactionConvector) ConvectorToRecentAddressTransactionResponse(res *client.ResGetRecentTransactionForAddress) *RecentAddressTransactionResponse {
	var items []TransactionItemRecent
	for _, item := range res.Items {
		items = append(items, TransactionItemRecent{
			BlockSignedAt:    item.BlockSignedAt,
			BlockHeight:      item.BlockHeight,
			BlockHash:        item.BlockHash,
			TxHash:           item.TxHash,
			TxOffset:         item.TxOffset,
			Successful:       item.Successful,
			FromAddress:      item.FromAddress,
			MinerAddress:     item.MinerAddress,
			FromAddressLabel: item.FromAddressLabel,
			ToAddress:        item.ToAddress,
			ToAddressLabel:   item.ToAddressLabel,
			Value:            item.Value,
			ValueQuote:       item.ValueQuote,
			PrettyValueQuote: item.PrettyValueQuote,
			GasMetadata: GasMetadataRecent{
				ContractDecimals:     item.GasMetadata.ContractDecimals,
				ContractName:         item.GasMetadata.ContractName,
				ContractTickerSymbol: item.GasMetadata.ContractTickerSymbol,
				ContractAddress:      item.GasMetadata.ContractAddress,
				SupportsERC:          item.GasMetadata.SupportsERC,
				LogoURL:              item.GasMetadata.LogoURL,
			},
			GasOffered:     item.GasOffered,
			GasSpent:       item.GasSpent,
			GasPrice:       item.GasPrice,
			FeesPaid:       item.FeesPaid,
			GasQuote:       item.GasQuote,
			PrettyGasQuote: item.PrettyGasQuote,
			GasQuoteRate:   item.GasQuoteRate,
			Explorers:      convertExplorers(item.Explorers),
			LogEvents:      convertLogEvents(item.LogEvents),
		})
	}

	return &RecentAddressTransactionResponse{
		Address:       res.Address,
		UpdatedAt:     res.UpdatedAt,
		QuoteCurrency: res.QuoteCurrency,
		ChainID:       res.ChainID,
		ChainName:     res.ChainName,
		CurrentPage:   res.CurrentPage,
		Links: Links{
			Prev: res.Links.Prev,
			Next: res.Links.Next,
		},
		Items: items,
	}
}

func convertExplorers(explorers []client.ExplorerRecent) []ExplorerRecent {
	var result []ExplorerRecent
	for _, explorer := range explorers {
		result = append(result, ExplorerRecent{
			Label: explorer.Label,
			URL:   explorer.URL,
		})
	}
	return result
}

func convertLogEvents(events []client.LogEventRecent) []LogEventRecent {
	var result []LogEventRecent
	for _, event := range events {
		var params []Param
		for _, p := range event.Decoded.Params {
			params = append(params, Param{
				Name:    p.Name,
				Type:    p.Type,
				Indexed: p.Indexed,
				Decoded: p.Decoded,
				Value:   p.Value,
			})
		}

		result = append(result, LogEventRecent{
			BlockSignedAt:          event.BlockSignedAt,
			BlockHeight:            event.BlockHeight,
			TxOffset:               event.TxOffset,
			LogOffset:              event.LogOffset,
			TxHash:                 event.TxHash,
			RawLogTopics:           event.RawLogTopics,
			SenderContractDecimals: event.SenderContractDecimals,
			SenderName:             event.SenderName,
			SenderTickerSymbol:     event.SenderTickerSymbol,
			SenderAddress:          event.SenderAddress,
			SenderAddressLabel:     event.SenderAddressLabel,
			SenderLogoURL:          event.SenderLogoURL,
			SupportsERC:            event.SupportsERC,
			SenderFactoryAddress:   event.SenderFactoryAddress,
			RawLogData:             event.RawLogData,
			Decoded: DecodedData{
				Name:      event.Decoded.Name,
				Signature: event.Decoded.Signature,
				Params:    params,
			},
		})
	}
	return result
}
