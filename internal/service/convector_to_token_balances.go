package service

import "goldrush-integration/pkg/integration/client"

type TokenBalancesConvector struct{}

func NewTokenBalancesConvector() *TokenBalancesConvector {
	return &TokenBalancesConvector{}
}

func (r *TokenBalancesConvector) ConvectorToTokenBalancesRequest(req *TokenBalancesRequest) *client.ReqGetTokenBalancesForAddress {
	return &client.ReqGetTokenBalancesForAddress{
		WalletAddress: req.WalletAddress,
		ChainName:     req.ChainName,
	}
}

func (r *TokenBalancesConvector) ConvectorToTokenBalancesResponse(res *client.ResGetTokenBalancesForAddress) *TokenBalancesResponse {
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
			QuoteRate:         item.QuoteRate,
			QuoteRate24h:      item.QuoteRate24h,
			Quote:             item.Quote,
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
