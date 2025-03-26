package service

import "github.com/vladislavprovich/goldrush-integration/pkg/client"

type ConvectorFromClient struct{}

func NewConvectorFromClient() *ConvectorFromClient {
	return &ConvectorFromClient{}
}

func (r *ConvectorFromClient) ConvectorToHistoricalPortfolioValueResponse(
	res *client.ResGetHistoricalPortfolioValueOverTime,
) *HistoricalPortfolioValueResponse {
	var items []HistoricalPortfolioItem
	for _, item := range res.Items {
		var holdings []HistoricalHolding
		for _, holding := range item.Holdings {
			holdings = append(holdings, HistoricalHolding{
				QuoteRate: holding.QuoteRate,
				Timestamp: holding.Timestamp,
				Close: HistoricalBalance{
					Balance:     holding.Close.Balance,
					Quote:       holding.Close.Quote,
					PrettyQuote: holding.Close.PrettyQuote,
				},
				High: HistoricalBalance{
					Balance:     holding.High.Balance,
					Quote:       holding.High.Quote,
					PrettyQuote: holding.High.PrettyQuote,
				},
				Low: HistoricalBalance{
					Balance:     holding.Low.Balance,
					Quote:       holding.Low.Quote,
					PrettyQuote: holding.Low.PrettyQuote,
				},
				Open: HistoricalBalance{
					Balance:     holding.Open.Balance,
					Quote:       holding.Open.Quote,
					PrettyQuote: holding.Open.PrettyQuote,
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

func (r *ConvectorFromClient) ConvectorToRecentAddressTransactionResponse(
	res *client.ResGetRecentTransactionForAddress,
) *RecentAddressTransactionResponse {
	var items []TransactionItemRecent
	for _, item := range res.Items {
		items = append(items, convertTransactionItemRecent(item))
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

func convertTransactionItemRecent(item client.TransactionItemRecent) TransactionItemRecent {
	return TransactionItemRecent{
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
		ValueQuote:       float64(item.ValueQuote),
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
		GasQuote:       float64(item.GasQuote),
		PrettyGasQuote: item.PrettyGasQuote,
		GasQuoteRate:   float64(item.GasQuoteRate),
		Explorers:      convertExplorers(item.Explorers),
		LogEvents:      convertLogEvents(item.LogEvents),
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
		if event.Decoded != nil {
			for _, p := range event.Decoded.Params {
				params = append(params, Param{
					Name:    p.Name,
					Type:    p.Type,
					Indexed: p.Indexed,
					Decoded: p.Decoded,
					Value:   p.Value,
				})
			}
		}

		decodedData := DecodedData{}
		if event.Decoded != nil {
			decodedData = DecodedData{
				Name:      event.Decoded.Name,
				Signature: event.Decoded.Signature,
				Params:    params,
			}
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
			Decoded:                decodedData,
		})
	}
	return result
}

func (r *ConvectorFromClient) ConvectorToTokenBalancesResponse(
	res *client.ResGetTokenBalancesForAddress,
) *TokenBalancesResponse {
	if res == nil {
		return nil
	}

	var items []TokenBalanceItem
	for _, item := range res.Items {
		items = append(items, TokenBalanceItem{
			ContractDecimals:     item.ContractDecimals,
			ContractName:         item.ContractName,
			ContractTickerSymbol: item.ContractTickerSymbol,
			ContractAddress:      item.ContractAddress,
			ContractDisplayName:  item.ContractDisplayName,
			SupportsERC:          item.SupportsERC,
			LogoURL:              item.LogoURL,
			LogoURLs: LogoURLs{
				TokenLogoURL:    item.LogoURLs.TokenLogoURL,
				ProtocolLogoURL: item.LogoURLs.ProtocolLogoURL,
				ChainLogoURL:    item.LogoURLs.ChainLogoURL,
			},
			LastTransferredAt: item.LastTransferredAt,
			NativeToken:       item.NativeToken,
			Type:              item.Type,
			IsSpam:            item.IsSpam,
			Balance:           item.Balance,
			Balance24h:        item.Balance24h,
			QuoteRate:         float64(item.QuoteRate),
			QuoteRate24h:      float64(item.QuoteRate24h),
			Quote:             float64(item.Quote),
			Quote24h:          item.Quote24h,
			PrettyQuote:       item.PrettyQuote,
			PrettyQuote24h:    item.PrettyQuote24h,
			ProtocolMetadata: ProtocolMetadata{
				ProtocolName: item.ProtocolMetadata.ProtocolName,
			},
			NFTData: convertNFTData(item.NFTData),
		})
	}

	return &TokenBalancesResponse{
		Address:       res.Address,
		ChainID:       res.ChainID,
		ChainName:     res.ChainName,
		QuoteCurrency: res.QuoteCurrency,
		UpdatedAt:     res.UpdatedAt,
		Items:         items,
	}
}

func convertNFTData(nftData []client.NFTData) []NFTData {
	var result []NFTData
	for _, data := range nftData {
		result = append(result, NFTData{
			TokenID:           data.TokenID,
			TokenBalance:      data.TokenBalance,
			TokenURL:          data.TokenURL,
			SupportsERC:       data.SupportsERC,
			TokenPriceWei:     data.TokenPriceWei,
			TokenQuoteRateEth: data.TokenQuoteRateEth,
			OriginalOwner:     data.OriginalOwner,
			ExternalData: ExternalData{
				Name:         data.ExternalData.Name,
				Description:  data.ExternalData.Description,
				Image:        data.ExternalData.Image,
				Image256:     data.ExternalData.Image256,
				Image512:     data.ExternalData.Image512,
				Image1024:    data.ExternalData.Image1024,
				AnimationURL: data.ExternalData.AnimationURL,
				ExternalURL:  data.ExternalData.ExternalURL,
				Owner:        data.ExternalData.Owner,
			},
			Owner:        data.Owner,
			OwnerAddress: data.OwnerAddress,
			Burned:       data.Burned,
		})
	}
	return result
}

func (r *ConvectorFromClient) ConvectorToTransactionInfoResponse(
	res *client.ResGetTransaction,
) *TransactionInfoResponse {
	// Check if the response is nil.
	if res == nil {
		return nil
	}

	// Create a response with data wrapper.
	response := &TransactionInfoResponse{
		ChainID:   int64(res.ChainID),
		ChainName: res.ChainName,
	}

	// Check if there are items in the response.
	if len(res.Items) == 0 {
		// Set basic fields for empty response.
		response.BlockSignedAt = res.UpdatedAt
		return response
	}

	tx := res.Items[0] // Get the first transaction item.
	response.BlockSignedAt = tx.BlockSignedAt
	response.BlockHeight = int64(tx.BlockHeight)
	response.BlockHash = tx.BlockHash
	response.TxHash = tx.TxHash
	response.TxOffset = int64(tx.TxOffset)
	response.Successful = tx.Successful
	response.FromAddress = tx.FromAddress
	response.MinerAddress = tx.MinerAddress
	response.FromAddressLabel = tx.FromAddressLabel
	response.ToAddress = tx.ToAddress
	response.ToAddressLabel = tx.ToAddressLabel
	response.Value = tx.Value
	response.ValueQuote = float64(tx.ValueQuote)
	response.PrettyValueQuote = tx.PrettyValueQuote
	response.GasMetadata = GasMetadataInfo{
		ContractDecimals:     int64(tx.GasMetadata.ContractDecimals),
		ContractName:         tx.GasMetadata.ContractName,
		ContractTickerSymbol: tx.GasMetadata.ContractTickerSymbol,
		ContractAddress:      tx.GasMetadata.ContractAddress,
		SupportsERC:          tx.GasMetadata.SupportsERC,
		LogoURL:              tx.GasMetadata.LogoURL,
	}
	response.GasOffered = int64(tx.GasOffered)
	response.GasSpent = int64(tx.GasSpent)
	response.GasPrice = int64(tx.GasPrice)
	response.FeesPaid = tx.FeesPaid
	response.GasQuote = float64(tx.GasQuote)
	response.PrettyGasQuote = tx.PrettyGasQuote
	response.GasQuoteRate = float64(tx.GasQuoteRate)
	response.Explorers = convertExplorersInfo(tx.Explorers)
	response.LogEvents = convertLogEventsInfo(tx.LogEvents)
	response.InternalTransfers = convertInternalTransfersInfo(tx.InternalTransfers)
	response.StateChanges = convertStateChangesInfo(tx.StateChanges)
	response.InputData = convertInputDataInfo(tx.InputData)

	return response
}

func convertInternalTransfersInfo(transfers []client.InternalTransferTransaction) []InternalTransferInfo {
	var result []InternalTransferInfo
	for _, transfer := range transfers {
		result = append(result, InternalTransferInfo{
			FromAddress: transfer.FromAddress,
			ToAddress:   transfer.ToAddress,
			Value:       transfer.Value,
			GasLimit:    int64(transfer.GasLimit),
		})
	}
	return result
}

func convertStateChangesInfo(changes []client.StateChangeTransaction) []StateChangeInfo {
	var result []StateChangeInfo
	for _, change := range changes {
		var storageChanges []StorageChangeInfo
		for _, storage := range change.StorageChanges {
			storageChanges = append(storageChanges, StorageChangeInfo{
				StorageAddress: storage.StorageAddress,
				ValueBefore:    storage.ValueBefore,
				ValueAfter:     storage.ValueAfter,
			})
		}

		result = append(result, StateChangeInfo{
			Address:        change.Address,
			BalanceBefore:  change.BalanceBefore,
			BalanceAfter:   change.BalanceAfter,
			StorageChanges: storageChanges,
			NonceBefore:    int64(change.NonceBefore),
			NonceAfter:     int64(change.NonceAfter),
		})
	}
	return result
}

func convertInputDataInfo(input client.InputDataTransaction) InputDataInfo {
	return InputDataInfo{
		MethodID: input.MethodID,
	}
}

func convertExplorersInfo(explorers []client.ExplorerTransaction) []ExplorerInfo {
	var result []ExplorerInfo
	for _, explorer := range explorers {
		result = append(result, ExplorerInfo{
			Label: explorer.Label,
			URL:   explorer.URL,
		})
	}
	return result
}

func convertLogEventsInfo(events []client.LogEventTransaction) []LogEventInfo {
	var result []LogEventInfo
	for _, event := range events {
		var params []ParamInfo
		for _, p := range event.Decoded.Params {
			params = append(params, ParamInfo{
				Name:    p.Name,
				Type:    p.Type,
				Indexed: p.Indexed,
				Decoded: p.Decoded,
				Value:   p.Value,
			})
		}

		result = append(result, LogEventInfo{
			BlockSignedAt:          event.BlockSignedAt,
			BlockHeight:            int64(event.BlockHeight),
			TxOffset:               int64(event.TxOffset),
			LogOffset:              int64(event.LogOffset),
			TxHash:                 event.TxHash,
			RawLogTopics:           event.RawLogTopics,
			SenderContractDecimals: int64(event.SenderContractDecimals),
			SenderName:             event.SenderName,
			SenderTickerSymbol:     event.SenderContractTickerSymbol,
			SenderAddress:          event.SenderAddress,
			SenderAddressLabel:     event.SenderAddressLabel,
			SenderLogoURL:          event.SenderLogoURL,
			SupportsERC:            event.SupportsERC,
			SenderFactoryAddress:   event.SenderFactoryAddress,
			RawLogData:             event.RawLogData,
			Decoded: DecodedDataInfo{
				Name:      event.Decoded.Name,
				Signature: event.Decoded.Signature,
				Params:    params,
			},
		})
	}
	return result
}
