package service

import "goldrush-integration/pkg/integration/client"

type TransactionInfoConvector struct{}

func NewTransactionInfoConvector() *TransactionInfoConvector {
	return &TransactionInfoConvector{}
}

func (r *TransactionInfoConvector) ConvectorToTransactionInfoRequest(req *TransactionInfoRequest) *client.ReqGetTransaction {
	return &client.ReqGetTransaction{
		ChainName: req.ChainName,
		TxHash:    req.TxHash,
	}
}

func (r *TransactionInfoConvector) ConvectorToTransactionInfoResponse(res *client.ResGetTransaction) *TransactionInfoResponse {
	if len(res.Items) == 0 {
		return nil
	}

	tx := res.Items[0] // Get the first transaction item
	return &TransactionInfoResponse{
		BlockSignedAt:    tx.BlockSignedAt,
		BlockHeight:      int64(tx.BlockHeight),
		BlockHash:        tx.BlockHash,
		TxHash:           tx.TxHash,
		TxOffset:         int64(tx.TxOffset),
		Successful:       tx.Successful,
		FromAddress:      tx.FromAddress,
		MinerAddress:     tx.MinerAddress,
		FromAddressLabel: tx.FromAddressLabel,
		ToAddress:        tx.ToAddress,
		ToAddressLabel:   tx.ToAddressLabel,
		Value:            tx.Value,
		ValueQuote:       float64(tx.ValueQuote),
		PrettyValueQuote: tx.PrettyValueQuote,
		GasMetadata: GasMetadataInfo{
			ContractDecimals:     int64(tx.GasMetadata.ContractDecimals),
			ContractName:         tx.GasMetadata.ContractName,
			ContractTickerSymbol: tx.GasMetadata.ContractTickerSymbol,
			ContractAddress:      tx.GasMetadata.ContractAddress,
			SupportsERC:          tx.GasMetadata.SupportsERC,
			LogoURL:              tx.GasMetadata.LogoURL,
		},
		GasOffered:        int64(tx.GasOffered),
		GasSpent:          int64(tx.GasSpent),
		GasPrice:          int64(tx.GasPrice),
		FeesPaid:          tx.FeesPaid,
		GasQuote:          float64(tx.GasQuote),
		PrettyGasQuote:    tx.PrettyGasQuote,
		GasQuoteRate:      float64(tx.GasQuoteRate),
		Explorers:         convertExplorersInfo(tx.Explorers),
		LogEvents:         convertLogEventsInfo(tx.LogEvents),
		InternalTransfers: convertInternalTransfersInfo(tx.InternalTransfers),
		StateChanges:      convertStateChangesInfo(tx.StateChanges),
		InputData:         convertInputDataInfo(tx.InputData),
	}
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
