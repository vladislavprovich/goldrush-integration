package service

import "goldrush-integration/pkg/integration/client"

type ConvectorToClient struct{}

func NewConvectorToClient() *ConvectorToClient {
	return &ConvectorToClient{}
}

func (r *ConvectorToClient) ConvectorToRecentAddressTransactionRequest(req *RecentAddressTransactionRequest) *client.ReqGetRecentTransactionForAddress {
	return &client.ReqGetRecentTransactionForAddress{
		WalletAddress: req.WalletAddress,
		ChainName:     req.ChainName,
	}
}

func (r *ConvectorToClient) ConvectorToHistoricalPortfolioValueRequest(req *HistoricalPortfolioValueRequest) *client.ReqGetHistoricalPortfolioValueOverTime {
	return &client.ReqGetHistoricalPortfolioValueOverTime{
		WalletAddress: req.WalletAddress,
		ChainName:     req.ChainName,
	}
}

func (r *ConvectorToClient) ConvectorToTokenBalancesRequest(req *TokenBalancesRequest) *client.ReqGetTokenBalancesForAddress {
	return &client.ReqGetTokenBalancesForAddress{
		WalletAddress: req.WalletAddress,
		ChainName:     req.ChainName,
	}
}

func (r *ConvectorToClient) ConvectorToTransactionInfoRequest(req *TransactionInfoRequest) *client.ReqGetTransaction {
	return &client.ReqGetTransaction{
		ChainName: req.ChainName,
		TxHash:    req.TxHash,
	}
}
