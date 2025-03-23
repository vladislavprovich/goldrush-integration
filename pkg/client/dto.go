package client

import "time"

type (
	ReqGetActivityAcrossAllChains struct {
		WalletAddress string `json:"wallet_address"`
		TestNets      bool   `json:"test_nets,omitempty" default:"false"`
	}

	RespGetActivityAcrossAllChains struct {
		UpdatedAt string          `json:"updated_at"`
		Address   string          `json:"address"`
		Items     []ChainActivity `json:"items"`
	}

	ChainActivity struct {
		Name          string `json:"name"`
		ChainID       string `json:"chain_id"`
		IsTestnet     bool   `json:"is_testnet"`
		Label         string `json:"label"`
		CategoryLabel string `json:"category_label"`
		PriorityLabel string `json:"priority_label"`
		LogoURL       string `json:"logo_url"`
		LastSeenAt    string `json:"last_seen_at"`
	}
)

type (
	ReqGetMultichainBalances struct {
		WalletAddress string `json:"wallet_address"`
	}

	ResGetMultichainBalances struct {
		UpdatedAt     time.Time     `json:"updated_at"`
		CursorBefore  string        `json:"cursor_before"`
		QuoteCurrency string        `json:"quote_currency"`
		Items         []BalanceItem `json:"items"`
	}

	BalanceItem struct {
		ContractDecimals    int              `json:"contract_decimals"`
		ContractName        string           `json:"contract_name"`
		ContractTicker      string           `json:"contract_ticker_symbol"`
		ContractAddress     string           `json:"contract_address"`
		ContractDisplayName string           `json:"contract_display_name"`
		SupportsERC         []string         `json:"supports_erc"`
		LogoURLs            LogoURLsBalances `json:"logo_urls"`
		LastTransferredAt   time.Time        `json:"last_transferred_at"`
		IsNativeToken       bool             `json:"is_native_token"`
		Type                string           `json:"type"`
		IsSpam              bool             `json:"is_spam"`
		Balance             string           `json:"balance"`
		Balance24h          string           `json:"balance_24h"`
		QuoteRate           float64          `json:"quote_rate"`
		QuoteRate24h        float64          `json:"quote_rate_24h"`
		Quote               float64          `json:"quote"`
		Quote24h            float64          `json:"quote_24h"`
		PrettyQuote         string           `json:"pretty_quote"`
		PrettyQuote24h      string           `json:"pretty_quote_24h"`
		ChainID             int              `json:"chain_id"`
		ChainName           string           `json:"chain_name"`
		ChainDisplayName    string           `json:"chain_display_name"`
	}

	LogoURLsBalances struct {
		TokenLogoURL    string `json:"token_logo_url"`
		ProtocolLogoURL string `json:"protocol_logo_url"`
		ChainLogoURL    string `json:"chain_logo_url"`
	}
)

type (
	ReqMultichainMultiaddressTransactions struct {
	}

	RespMultichainMultiaddressTransactions struct {
		UpdatedAt     time.Time         `json:"updated_at"`
		CursorBefore  string            `json:"cursor_before"`
		CursorAfter   string            `json:"cursor_after"`
		QuoteCurrency string            `json:"quote_currency"`
		Items         []TransactionItem `json:"items"`
	}

	TransactionItem struct {
		BlockHeight       int                `json:"block_height"`
		BlockSignedAt     time.Time          `json:"block_signed_at"`
		BlockHash         string             `json:"block_hash"`
		TxHash            string             `json:"tx_hash"`
		TxOffset          int                `json:"tx_offset"`
		MinerAddress      string             `json:"miner_address"`
		FromAddress       string             `json:"from_address"`
		FromLabel         string             `json:"from_address_label,omitempty"`
		ToAddress         string             `json:"to_address"`
		ToLabel           string             `json:"to_address_label,omitempty"`
		Value             string             `json:"value"`
		ValueQuote        float64            `json:"value_quote"`
		PrettyValueQuote  string             `json:"pretty_value_quote,omitempty"`
		GasOffered        int                `json:"gas_offered"`
		GasSpent          int                `json:"gas_spent"`
		GasPrice          int                `json:"gas_price"`
		GasQuote          float64            `json:"gas_quote"`
		PrettyGasQuote    string             `json:"pretty_gas_quote,omitempty"`
		GasQuoteRate      float64            `json:"gas_quote_rate"`
		FeesPaid          string             `json:"fees_paid"`
		GasMetadata       GasMetadata        `json:"gas_metadata"`
		Successful        bool               `json:"successful"`
		ChainID           string             `json:"chain_id"`
		ChainName         string             `json:"chain_name"`
		Explorers         []Explorer         `json:"explorers"`
		LogEvents         []LogEvent         `json:"log_events"`
		InternalTransfers []InternalTransfer `json:"internal_transfers,omitempty"`
		StateChanges      []StateChange      `json:"state_changes,omitempty"`
		InputData         InputData          `json:"input_data,omitempty"`
	}

	GasMetadata struct {
		ContractDecimals int      `json:"contract_decimals"`
		ContractName     string   `json:"contract_name"`
		ContractSymbol   string   `json:"contract_ticker_symbol"`
		ContractAddress  string   `json:"contract_address"`
		SupportsERC      []string `json:"supports_erc,omitempty"`
		LogoURL          string   `json:"logo_url"`
	}

	Explorer struct {
		Label string `json:"label"`
		URL   string `json:"url"`
	}

	LogEvent struct {
		BlockSignedAt          time.Time       `json:"block_signed_at"`
		BlockHeight            int             `json:"block_height"`
		TxOffset               int             `json:"tx_offset"`
		LogOffset              int             `json:"log_offset"`
		TxHash                 string          `json:"tx_hash"`
		RawLogTopics           []string        `json:"raw_log_topics"`
		SenderContractDecimals int             `json:"sender_contract_decimals"`
		SenderName             string          `json:"sender_name"`
		SenderSymbol           string          `json:"sender_contract_ticker_symbol"`
		SenderAddress          string          `json:"sender_address"`
		SenderLabel            string          `json:"sender_address_label,omitempty"`
		SenderLogoURL          string          `json:"sender_logo_url"`
		SupportsERC            []string        `json:"supports_erc,omitempty"`
		SenderFactoryAddress   string          `json:"sender_factory_address"`
		RawLogData             string          `json:"raw_log_data"`
		Decoded                DecodedLogEvent `json:"decoded"`
	}

	DecodedLogEvent struct {
		Name      string     `json:"name"`
		Signature string     `json:"signature"`
		Params    []LogParam `json:"params"`
	}

	LogParam struct {
		Name    string `json:"name"`
		Type    string `json:"type"`
		Indexed bool   `json:"indexed"`
		Decoded bool   `json:"decoded"`
		Value   string `json:"value"`
	}

	InternalTransfer struct {
		FromAddress string `json:"from_address"`
		ToAddress   string `json:"to_address"`
		Value       string `json:"value"`
		GasLimit    int    `json:"gas_limit"`
	}

	StateChange struct {
		Address        string          `json:"address"`
		BalanceBefore  string          `json:"balance_before"`
		BalanceAfter   string          `json:"balance_after"`
		StorageChanges []StorageChange `json:"storage_changes,omitempty"`
		NonceBefore    int             `json:"nonce_before"`
		NonceAfter     int             `json:"nonce_after"`
	}

	StorageChange struct {
		StorageAddress string `json:"storage_address"`
		ValueBefore    string `json:"value_before"`
		ValueAfter     string `json:"value_after"`
	}

	InputData struct {
		MethodID string `json:"method_id"`
	}
)

type (
	ReqGetTokenBalancesForAddress struct {
		WalletAddress string `json:"wallet_address"`
		ChainName     string `json:"chain_name"`
	}

	ResGetTokenBalancesForAddress struct {
		Address       string             `json:"address"`
		ChainID       int                `json:"chain_id"`
		ChainName     string             `json:"chain_name"`
		QuoteCurrency string             `json:"quote_currency"`
		UpdatedAt     string             `json:"updated_at"`
		Items         []TokenBalanceItem `json:"items"`
	}

	LogoURLs struct {
		TokenLogoURL    string `json:"token_logo_url"`
		ProtocolLogoURL string `json:"protocol_logo_url"`
		ChainLogoURL    string `json:"chain_logo_url"`
	}

	ProtocolMetadata struct {
		ProtocolName string `json:"protocol_name"`
	}

	ExternalData struct {
		Name         string `json:"name"`
		Description  string `json:"description"`
		Image        string `json:"image"`
		Image256     string `json:"image_256"`
		Image512     string `json:"image_512"`
		Image1024    string `json:"image_1024"`
		AnimationURL string `json:"animation_url"`
		ExternalURL  string `json:"external_url"`
		Owner        string `json:"owner"`
	}

	NFTData struct {
		TokenID           string       `json:"token_id"`
		TokenBalance      string       `json:"token_balance"`
		TokenURL          string       `json:"token_url"`
		SupportsERC       []string     `json:"supports_erc"`
		TokenPriceWei     string       `json:"token_price_wei"`
		TokenQuoteRateEth string       `json:"token_quote_rate_eth"`
		OriginalOwner     string       `json:"original_owner"`
		ExternalData      ExternalData `json:"external_data"`
		Owner             string       `json:"owner"`
		OwnerAddress      string       `json:"owner_address"`
		Burned            bool         `json:"burned"`
	}

	TokenBalanceItem struct {
		ContractDecimals     int              `json:"contract_decimals"`
		ContractName         string           `json:"contract_name"`
		ContractTickerSymbol string           `json:"contract_ticker_symbol"`
		ContractAddress      string           `json:"contract_address"`
		ContractDisplayName  string           `json:"contract_display_name"`
		SupportsERC          []string         `json:"supports_erc"`
		LogoURL              string           `json:"logo_url"`
		LogoURLs             LogoURLs         `json:"logo_urls"`
		LastTransferredAt    string           `json:"last_transferred_at"`
		NativeToken          bool             `json:"native_token"`
		Type                 string           `json:"type"`
		IsSpam               bool             `json:"is_spam"`
		Balance              string           `json:"balance"`
		Balance24h           string           `json:"balance_24h"`
		QuoteRate            int              `json:"quote_rate"`
		QuoteRate24h         int              `json:"quote_rate_24h"`
		Quote                int              `json:"quote"`
		Quote24h             int              `json:"quote_24h"`
		PrettyQuote          string           `json:"pretty_quote"`
		PrettyQuote24h       string           `json:"pretty_quote_24h"`
		ProtocolMetadata     ProtocolMetadata `json:"protocol_metadata"`
		NFTData              []NFTData        `json:"nft_data"`
	}
)

type (
	ReqGetHistoricalTokenBalancesForAddress struct {
		WalletAddress string `json:"wallet_address"`
		ChainName     string `json:"chain_name"`
	}

	ResGetHistoricalTokenBalancesForAddress struct {
		Address       string             `json:"address"`
		UpdatedAt     string             `json:"updated_at"`
		QuoteCurrency string             `json:"quote_currency"`
		ChainID       int                `json:"chain_id"`
		ChainName     string             `json:"chain_name"`
		Items         []TokenBalanceItem `json:"items"`
	}

	ProtocolMetadataHistory struct {
		ProtocolName string `json:"protocol_name"`
	}

	ExternalDataHistory struct {
		Name         string `json:"name"`
		Description  string `json:"description"`
		Image        string `json:"image"`
		Image256     string `json:"image_256"`
		Image512     string `json:"image_512"`
		Image1024    string `json:"image_1024"`
		AnimationURL string `json:"animation_url"`
		ExternalURL  string `json:"external_url"`
		Owner        string `json:"owner"`
	}

	NFTAttribute struct {
		TraitType string `json:"trait_type"`
		Value     string `json:"value"`
	}

	NFTThumbnails struct {
		Image256          string `json:"image_256"`
		Image512          string `json:"image_512"`
		Image1024         string `json:"image_1024"`
		ImageOpenGraphURL string `json:"image_opengraph_url"`
		ThumbHash         string `json:"thumbhash"`
	}

	AssetProperties struct {
		AssetWidth    int    `json:"asset_width"`
		AssetHeight   int    `json:"asset_height"`
		DominantColor string `json:"dominant_color"`
	}

	NFTDataHistory struct {
		TokenID           string              `json:"token_id"`
		TokenBalance      string              `json:"token_balance"`
		TokenURL          string              `json:"token_url"`
		SupportsERC       []string            `json:"supports_erc"`
		TokenPriceWei     string              `json:"token_price_wei"`
		TokenQuoteRateEth string              `json:"token_quote_rate_eth"`
		OriginalOwner     string              `json:"original_owner"`
		ExternalData      ExternalDataHistory `json:"external_data"`
		Owner             string              `json:"owner"`
		OwnerAddress      string              `json:"owner_address"`
		Burned            bool                `json:"burned"`
		Attributes        []NFTAttribute      `json:"attributes"`
		Thumbnails        NFTThumbnails       `json:"thumbnails"`
		AssetProperties   AssetProperties     `json:"asset_properties"`
	}

	TokenBalanceItemHistory struct {
		ContractDecimals           int                     `json:"contract_decimals"`
		ContractName               string                  `json:"contract_name"`
		ContractTickerSymbol       string                  `json:"contract_ticker_symbol"`
		ContractAddress            string                  `json:"contract_address"`
		SupportsERC                []string                `json:"supports_erc"`
		LogoURL                    string                  `json:"logo_url"`
		BlockHeight                int                     `json:"block_height"`
		LastTransferredBlockHeight int                     `json:"last_transferred_block_height"`
		ContractDisplayName        string                  `json:"contract_display_name"`
		LastTransferredAt          string                  `json:"last_transferred_at"`
		NativeToken                bool                    `json:"native_token"`
		Type                       string                  `json:"type"`
		IsSpam                     bool                    `json:"is_spam"`
		Balance                    string                  `json:"balance"`
		QuoteRate                  int                     `json:"quote_rate"`
		Quote                      int                     `json:"quote"`
		PrettyQuote                string                  `json:"pretty_quote"`
		ProtocolMetadata           ProtocolMetadataHistory `json:"protocol_metadata"`
		NFTData                    []NFTDataHistory        `json:"nft_data"`
	}
)

type (
	ReqGetNativeTokenBalancesForAddress struct {
		WalletAddress string `json:"wallet_address"`
		ChainName     string `json:"chain_name"`
	}

	ResGetNativeTokenBalancesForAddress struct {
		Address       string                   `json:"address"`
		UpdatedAt     string                   `json:"updated_at"`
		QuoteCurrency string                   `json:"quote_currency"`
		ChainID       int                      `json:"chain_id"`
		ChainName     string                   `json:"chain_name"`
		Items         []TokenBalanceItemNative `json:"items"`
	}

	TokenBalanceItemNative struct {
		ContractDecimals     int      `json:"contract_decimals"`
		ContractName         string   `json:"contract_name"`
		ContractTickerSymbol string   `json:"contract_ticker_symbol"`
		ContractAddress      string   `json:"contract_address"`
		SupportsERC          []string `json:"supports_erc"`
		LogoURL              string   `json:"logo_url"`
		BlockHeight          int      `json:"block_height"`
		Balance              string   `json:"balance"`
		QuoteRate            int      `json:"quote_rate"`
		Quote                int      `json:"quote"`
		PrettyQuote          string   `json:"pretty_quote"`
	}
)

type (
	ReqGetERC20TokenTransfersForAddress struct {
		WalletAddress string `json:"wallet_address"`
		ChainName     string `json:"chain_name"`
	}

	ResGetERC20TokenTransfersForAddress struct {
		Address       string              `json:"address"`
		UpdatedAt     string              `json:"updated_at"`
		QuoteCurrency string              `json:"quote_currency"`
		ChainID       int                 `json:"chain_id"`
		ChainName     string              `json:"chain_name"`
		Items         []TokenTransferItem `json:"items"`
		Pagination    Pagination          `json:"pagination"`
	}

	GasMetadataERC20 struct {
		ContractDecimals     int      `json:"contract_decimals"`
		ContractName         string   `json:"contract_name"`
		ContractTickerSymbol string   `json:"contract_ticker_symbol"`
		ContractAddress      string   `json:"contract_address"`
		SupportsERC          []string `json:"supports_erc"`
		LogoURL              string   `json:"logo_url"`
	}

	MethodCall struct {
		SenderAddress string `json:"sender_address"`
		Method        string `json:"method"`
	}

	ExplorerERC20 struct {
		Label string `json:"label"`
		URL   string `json:"url"`
	}

	Transfer struct {
		BlockSignedAt        string          `json:"block_signed_at"`
		TxHash               string          `json:"tx_hash"`
		FromAddress          string          `json:"from_address"`
		FromAddressLabel     string          `json:"from_address_label"`
		ToAddress            string          `json:"to_address"`
		ToAddressLabel       string          `json:"to_address_label"`
		ContractDecimals     int             `json:"contract_decimals"`
		ContractName         string          `json:"contract_name"`
		ContractTickerSymbol string          `json:"contract_ticker_symbol"`
		ContractAddress      string          `json:"contract_address"`
		LogoURL              string          `json:"logo_url"`
		TransferType         string          `json:"transfer_type"`
		Delta                string          `json:"delta"`
		Balance              string          `json:"balance"`
		QuoteRate            int             `json:"quote_rate"`
		DeltaQuote           int             `json:"delta_quote"`
		PrettyDeltaQuote     string          `json:"pretty_delta_quote"`
		BalanceQuote         int             `json:"balance_quote"`
		MethodCalls          []MethodCall    `json:"method_calls"`
		Explorers            []ExplorerERC20 `json:"explorers"`
	}

	TokenTransferItem struct {
		BlockSignedAt  string           `json:"block_signed_at"`
		BlockHeight    int              `json:"block_height"`
		BlockHash      string           `json:"block_hash"`
		TxHash         string           `json:"tx_hash"`
		TxOffset       int              `json:"tx_offset"`
		Successful     bool             `json:"successful"`
		MinerAddress   string           `json:"miner_address"`
		FromAddress    string           `json:"from_address"`
		FromLabel      string           `json:"from_address_label"`
		ToAddress      string           `json:"to_address"`
		ToLabel        string           `json:"to_address_label"`
		Value          string           `json:"value"`
		ValueQuote     int              `json:"value_quote"`
		PrettyValue    string           `json:"pretty_value_quote"`
		GasMetadata    GasMetadataERC20 `json:"gas_metadata"`
		GasOffered     int              `json:"gas_offered"`
		GasSpent       int              `json:"gas_spent"`
		GasPrice       int              `json:"gas_price"`
		FeesPaid       string           `json:"fees_paid"`
		GasQuote       int              `json:"gas_quote"`
		PrettyGasQuote string           `json:"pretty_gas_quote"`
		GasQuoteRate   int              `json:"gas_quote_rate"`
		Transfers      []Transfer       `json:"transfers"`
	}

	Pagination struct {
		HasMore    bool `json:"has_more"`
		PageNumber int  `json:"page_number"`
		PageSize   int  `json:"page_size"`
		TotalCount int  `json:"total_count"`
	}
)

type (
	ReqGetHistoricalPortfolioValueOverTime struct {
		WalletAddress string `json:"wallet_address"`
		ChainName     string `json:"chain_name"`
	}

	PricePoint struct {
		Balance     string `json:"balance"`
		Quote       int    `json:"quote"`
		PrettyQuote string `json:"pretty_quote"`
	}

	Holding struct {
		QuoteRate int        `json:"quote_rate"`
		Timestamp string     `json:"timestamp"`
		Close     PricePoint `json:"close"`
		High      PricePoint `json:"high"`
		Low       PricePoint `json:"low"`
		Open      PricePoint `json:"open"`
	}

	HistoricalPortfolioItem struct {
		ContractAddress      string    `json:"contract_address"`
		ContractDecimals     int       `json:"contract_decimals"`
		ContractName         string    `json:"contract_name"`
		ContractTickerSymbol string    `json:"contract_ticker_symbol"`
		LogoURL              string    `json:"logo_url"`
		Holdings             []Holding `json:"holdings"`
	}

	ResGetHistoricalPortfolioValueOverTime struct {
		Address       string                    `json:"address"`
		UpdatedAt     string                    `json:"updated_at"`
		QuoteCurrency string                    `json:"quote_currency"`
		ChainID       int                       `json:"chain_id"`
		ChainName     string                    `json:"chain_name"`
		Items         []HistoricalPortfolioItem `json:"items"`
	}
)

type (
	ReqGetTokenHoldersBlockHeight struct {
		WalletAddress string `json:"wallet_address"`
		ChainName     string `json:"chain_name"`
	}

	ResGetTokenHoldersBlockHeight struct {
		UpdatedAt  string                `json:"updated_at"`
		ChainID    int                   `json:"chain_id"`
		ChainName  string                `json:"chain_name"`
		Items      []TokenHolderItem     `json:"items"`
		Pagination PaginationBlockHeight `json:"pagination"`
	}

	TokenHolderItem struct {
		ContractDecimals     int      `json:"contract_decimals"`
		ContractName         string   `json:"contract_name"`
		ContractTickerSymbol string   `json:"contract_ticker_symbol"`
		ContractAddress      string   `json:"contract_address"`
		SupportsERC          []string `json:"supports_erc"`
		LogoURL              string   `json:"logo_url"`
		Address              string   `json:"address"`
		Balance              string   `json:"balance"`
		TotalSupply          string   `json:"total_supply"`
		BlockHeight          int      `json:"block_height"`
	}

	PaginationBlockHeight struct {
		HasMore    bool `json:"has_more"`
		PageNumber int  `json:"page_number"`
		PageSize   int  `json:"page_size"`
		TotalCount int  `json:"total_count"`
	}
)

type (
	ReqGetTransaction struct {
		ChainName string `json:"chainName"`
		TxHash    string `json:"txHash"`
	}

	ResGetTransaction struct {
		UpdatedAt string                       `json:"updated_at"`
		ChainID   int                          `json:"chain_id"`
		ChainName string                       `json:"chain_name"`
		Items     []TransactionItemTransaction `json:"items"`
	}

	GasMetadataTransaction struct {
		ContractDecimals     int      `json:"contract_decimals"`
		ContractName         string   `json:"contract_name"`
		ContractTickerSymbol string   `json:"contract_ticker_symbol"`
		ContractAddress      string   `json:"contract_address"`
		SupportsERC          []string `json:"supports_erc"`
		LogoURL              string   `json:"logo_url"`
	}

	ExplorerTransaction struct {
		Label string `json:"label"`
		URL   string `json:"url"`
	}

	DecodedParam struct {
		Name    string `json:"name"`
		Type    string `json:"type"`
		Indexed bool   `json:"indexed"`
		Decoded bool   `json:"decoded"`
		Value   string `json:"value"`
	}

	DecodedLog struct {
		Name      string         `json:"name"`
		Signature string         `json:"signature"`
		Params    []DecodedParam `json:"params"`
	}

	LogEventTransaction struct {
		BlockSignedAt              string     `json:"block_signed_at"`
		BlockHeight                int        `json:"block_height"`
		TxOffset                   int        `json:"tx_offset"`
		LogOffset                  int        `json:"log_offset"`
		TxHash                     string     `json:"tx_hash"`
		RawLogTopics               []string   `json:"raw_log_topics"`
		SenderContractDecimals     int        `json:"sender_contract_decimals"`
		SenderName                 string     `json:"sender_name"`
		SenderContractTickerSymbol string     `json:"sender_contract_ticker_symbol"`
		SenderAddress              string     `json:"sender_address"`
		SenderAddressLabel         string     `json:"sender_address_label"`
		SenderLogoURL              string     `json:"sender_logo_url"`
		SupportsERC                []string   `json:"supports_erc"`
		SenderFactoryAddress       string     `json:"sender_factory_address"`
		RawLogData                 string     `json:"raw_log_data"`
		Decoded                    DecodedLog `json:"decoded"`
	}

	InternalTransferTransaction struct {
		FromAddress string `json:"from_address"`
		ToAddress   string `json:"to_address"`
		Value       string `json:"value"`
		GasLimit    int    `json:"gas_limit"`
	}

	StorageChangeTransaction struct {
		StorageAddress string `json:"storage_address"`
		ValueBefore    string `json:"value_before"`
		ValueAfter     string `json:"value_after"`
	}

	StateChangeTransaction struct {
		Address        string                     `json:"address"`
		BalanceBefore  string                     `json:"balance_before"`
		BalanceAfter   string                     `json:"balance_after"`
		StorageChanges []StorageChangeTransaction `json:"storage_changes"`
		NonceBefore    int                        `json:"nonce_before"`
		NonceAfter     int                        `json:"nonce_after"`
	}

	InputDataTransaction struct {
		MethodID string `json:"method_id"`
	}

	TransactionItemTransaction struct {
		BlockSignedAt     string                        `json:"block_signed_at"`
		BlockHeight       int                           `json:"block_height"`
		BlockHash         string                        `json:"block_hash"`
		TxHash            string                        `json:"tx_hash"`
		TxOffset          int                           `json:"tx_offset"`
		Successful        bool                          `json:"successful"`
		FromAddress       string                        `json:"from_address"`
		MinerAddress      string                        `json:"miner_address"`
		FromAddressLabel  string                        `json:"from_address_label"`
		ToAddress         string                        `json:"to_address"`
		ToAddressLabel    string                        `json:"to_address_label"`
		Value             string                        `json:"value"`
		ValueQuote        int                           `json:"value_quote"`
		PrettyValueQuote  string                        `json:"pretty_value_quote"`
		GasMetadata       GasMetadataTransaction        `json:"gas_metadata"`
		GasOffered        int                           `json:"gas_offered"`
		GasSpent          int                           `json:"gas_spent"`
		GasPrice          int                           `json:"gas_price"`
		FeesPaid          string                        `json:"fees_paid"`
		GasQuote          int                           `json:"gas_quote"`
		PrettyGasQuote    string                        `json:"pretty_gas_quote"`
		GasQuoteRate      int                           `json:"gas_quote_rate"`
		Explorers         []ExplorerTransaction         `json:"explorers"`
		LogEvents         []LogEventTransaction         `json:"log_events"`
		InternalTransfers []InternalTransferTransaction `json:"internal_transfers"`
		StateChanges      []StateChangeTransaction      `json:"state_changes"`
		InputData         InputDataTransaction          `json:"input_data"`
	}
)

type (
	ReqGetTransactionSummaryForAddress struct {
		WalletAddress string `json:"wallet_address"`
		ChainName     string `json:"chain_name"`
	}

	ResGetTransactionSummaryForAddress struct {
		UpdatedAt string                   `json:"updated_at"`
		Address   string                   `json:"address"`
		ChainID   int                      `json:"chain_id"`
		ChainName string                   `json:"chain_name"`
		Items     []TransactionSummaryItem `json:"items"`
	}

	TransactionDetail struct {
		BlockSignedAt string `json:"block_signed_at"`
		TxHash        string `json:"tx_hash"`
		TxDetailLink  string `json:"tx_detail_link"`
	}

	GasMetadataSummaryForAddress struct {
		ContractDecimals     int      `json:"contract_decimals"`
		ContractName         string   `json:"contract_name"`
		ContractTickerSymbol string   `json:"contract_ticker_symbol"`
		ContractAddress      string   `json:"contract_address"`
		SupportsERC          []string `json:"supports_erc"`
		LogoURL              string   `json:"logo_url"`
	}

	GasSummary struct {
		TotalSentCount             int                          `json:"total_sent_count"`
		TotalFeesPaid              string                       `json:"total_fees_paid"`
		TotalGasQuote              int                          `json:"total_gas_quote"`
		PrettyTotalGasQuote        string                       `json:"pretty_total_gas_quote"`
		AverageGasQuotePerTx       int                          `json:"average_gas_quote_per_tx"`
		PrettyAverageGasQuotePerTx string                       `json:"pretty_average_gas_quote_per_tx"`
		GasMetadata                GasMetadataSummaryForAddress `json:"gas_metadata"`
	}

	TransactionSummaryItem struct {
		TotalCount          int               `json:"total_count"`
		EarliestTransaction TransactionDetail `json:"earliest_transaction"`
		LatestTransaction   TransactionDetail `json:"latest_transaction"`
		GasSummary          GasSummary        `json:"gas_summary"`
	}
)

type (
	ReqGetEarliesrTransactionForAddress struct {
		WalletAddress string `json:"wallet_address"`
		ChainName     string `json:"chain_name"`
	}

	ResGetEarliestTransactionForAddress struct {
		Address       string                    `json:"address"`
		UpdatedAt     string                    `json:"updated_at"`
		QuoteCurrency string                    `json:"quote_currency"`
		ChainID       int                       `json:"chain_id"`
		ChainName     string                    `json:"chain_name"`
		Complete      bool                      `json:"complete"`
		CurrentBucket int                       `json:"current_bucket"`
		Links         Link                      `json:"links"`
		Items         []TransactionItemEarliest `json:"items"`
	}

	Link struct {
		Prev string `json:"prev"`
		Next string `json:"next"`
	}

	ExplorerEarliest struct {
		Label string `json:"label"`
		URL   string `json:"url"`
	}

	GasMetadataEarliest struct {
		ContractDecimals     int      `json:"contract_decimals"`
		ContractName         string   `json:"contract_name"`
		ContractTickerSymbol string   `json:"contract_ticker_symbol"`
		ContractAddress      string   `json:"contract_address"`
		SupportsERC          []string `json:"supports_erc"`
		LogoURL              string   `json:"logo_url"`
	}

	LogParamEarliest struct {
		Name    string `json:"name"`
		Type    string `json:"type"`
		Indexed bool   `json:"indexed"`
		Decoded bool   `json:"decoded"`
		Value   string `json:"value"`
	}

	DecodedLogEarliest struct {
		Name      string             `json:"name"`
		Signature string             `json:"signature"`
		Params    []LogParamEarliest `json:"params"`
	}

	LogEventEarliest struct {
		BlockSignedAt              string             `json:"block_signed_at"`
		BlockHeight                int                `json:"block_height"`
		TxOffset                   int                `json:"tx_offset"`
		LogOffset                  int                `json:"log_offset"`
		TxHash                     string             `json:"tx_hash"`
		RawLogTopics               []string           `json:"raw_log_topics"`
		SenderContractDecimals     int                `json:"sender_contract_decimals"`
		SenderName                 string             `json:"sender_name"`
		SenderContractTickerSymbol string             `json:"sender_contract_ticker_symbol"`
		SenderAddress              string             `json:"sender_address"`
		SenderAddressLabel         string             `json:"sender_address_label"`
		SenderLogoURL              string             `json:"sender_logo_url"`
		SupportsERC                []string           `json:"supports_erc"`
		SenderFactoryAddress       string             `json:"sender_factory_address"`
		RawLogData                 string             `json:"raw_log_data"`
		Decoded                    DecodedLogEarliest `json:"decoded"`
	}

	TransactionItemEarliest struct {
		BlockSignedAt    string              `json:"block_signed_at"`
		BlockHeight      int                 `json:"block_height"`
		BlockHash        string              `json:"block_hash"`
		TxHash           string              `json:"tx_hash"`
		TxOffset         int                 `json:"tx_offset"`
		Successful       bool                `json:"successful"`
		FromAddress      string              `json:"from_address"`
		MinerAddress     string              `json:"miner_address"`
		FromAddressLabel string              `json:"from_address_label"`
		ToAddress        string              `json:"to_address"`
		ToAddressLabel   string              `json:"to_address_label"`
		Value            string              `json:"value"`
		ValueQuote       int                 `json:"value_quote"`
		PrettyValueQuote string              `json:"pretty_value_quote"`
		GasMetadata      GasMetadataEarliest `json:"gas_metadata"`
		GasOffered       int                 `json:"gas_offered"`
		GasSpent         int                 `json:"gas_spent"`
		GasPrice         int                 `json:"gas_price"`
		FeesPaid         string              `json:"fees_paid"`
		GasQuote         int                 `json:"gas_quote"`
		PrettyGasQuote   string              `json:"pretty_gas_quote"`
		GasQuoteRate     int                 `json:"gas_quote_rate"`
		Explorers        []ExplorerEarliest  `json:"explorers"`
		LogEvents        []LogEventEarliest  `json:"log_events"`
	}
)

type (
	ReqGetRecentTransactionForAddress struct {
		WalletAddress string `json:"wallet_address"`
		ChainName     string `json:"chain_name"`
	}

	ResGetRecentTransactionForAddress struct {
		Address       string                  `json:"address"`
		UpdatedAt     string                  `json:"updated_at"`
		QuoteCurrency string                  `json:"quote_currency"`
		ChainID       int                     `json:"chain_id"`
		ChainName     string                  `json:"chain_name"`
		CurrentPage   int                     `json:"current_page"`
		Links         Links                   `json:"links"`
		Items         []TransactionItemRecent `json:"items"`
	}

	Links struct {
		Prev string `json:"prev"`
		Next string `json:"next"`
	}

	TransactionItemRecent struct {
		BlockSignedAt    string            `json:"block_signed_at"`
		BlockHeight      int               `json:"block_height"`
		BlockHash        string            `json:"block_hash"`
		TxHash           string            `json:"tx_hash"`
		TxOffset         int               `json:"tx_offset"`
		Successful       bool              `json:"successful"`
		FromAddress      string            `json:"from_address"`
		MinerAddress     string            `json:"miner_address"`
		FromAddressLabel string            `json:"from_address_label"`
		ToAddress        string            `json:"to_address"`
		ToAddressLabel   string            `json:"to_address_label"`
		Value            string            `json:"value"`
		ValueQuote       int               `json:"value_quote"`
		PrettyValueQuote string            `json:"pretty_value_quote"`
		GasMetadata      GasMetadataRecent `json:"gas_metadata"`
		GasOffered       int               `json:"gas_offered"`
		GasSpent         int               `json:"gas_spent"`
		GasPrice         int               `json:"gas_price"`
		FeesPaid         string            `json:"fees_paid"`
		GasQuote         int               `json:"gas_quote"`
		PrettyGasQuote   string            `json:"pretty_gas_quote"`
		GasQuoteRate     int               `json:"gas_quote_rate"`
		Explorers        []ExplorerRecent  `json:"explorers"`
		LogEvents        []LogEventRecent  `json:"log_events"`
	}

	GasMetadataRecent struct {
		ContractDecimals     int      `json:"contract_decimals"`
		ContractName         string   `json:"contract_name"`
		ContractTickerSymbol string   `json:"contract_ticker_symbol"`
		ContractAddress      string   `json:"contract_address"`
		SupportsERC          []string `json:"supports_erc"`
		LogoURL              string   `json:"logo_url"`
	}

	ExplorerRecent struct {
		Label string `json:"label"`
		URL   string `json:"url"`
	}

	LogEventRecent struct {
		BlockSignedAt          string      `json:"block_signed_at"`
		BlockHeight            int         `json:"block_height"`
		TxOffset               int         `json:"tx_offset"`
		LogOffset              int         `json:"log_offset"`
		TxHash                 string      `json:"tx_hash"`
		RawLogTopics           []string    `json:"raw_log_topics"`
		SenderContractDecimals int         `json:"sender_contract_decimals"`
		SenderName             string      `json:"sender_name"`
		SenderTickerSymbol     string      `json:"sender_contract_ticker_symbol"`
		SenderAddress          string      `json:"sender_address"`
		SenderAddressLabel     string      `json:"sender_address_label"`
		SenderLogoURL          string      `json:"sender_logo_url"`
		SupportsERC            []string    `json:"supports_erc"`
		SenderFactoryAddress   string      `json:"sender_factory_address"`
		RawLogData             string      `json:"raw_log_data"`
		Decoded                DecodedData `json:"decoded"`
	}

	DecodedData struct {
		Name      string  `json:"name"`
		Signature string  `json:"signature"`
		Params    []Param `json:"params"`
	}

	Param struct {
		Name    string `json:"name"`
		Type    string `json:"type"`
		Indexed bool   `json:"indexed"`
		Decoded bool   `json:"decoded"`
		Value   string `json:"value"`
	}
)

type (
	ReqGetPaginatedTransactionForAddress struct {
		WalletAddress string `json:"wallet_address"`
		ChainName     string `json:"chain_name"`
		Page          int    `json:"page"`
	}

	ResGetPaginatedTransactionForAddress struct {
		Address       string         `json:"address"`
		UpdatedAt     string         `json:"updated_at"`
		QuoteCurrency string         `json:"quote_currency"`
		ChainID       int            `json:"chain_id"`
		ChainName     string         `json:"chain_name"`
		CurrentPage   int            `json:"current_page"`
		Links         LinksPaginated `json:"links"`
		Items         []Item         `json:"items"`
	}

	LinksPaginated struct {
		Prev string `json:"prev"`
		Next string `json:"next"`
	}

	Item struct {
		BlockSignedAt    string               `json:"block_signed_at"`
		BlockHeight      int                  `json:"block_height"`
		BlockHash        string               `json:"block_hash"`
		TxHash           string               `json:"tx_hash"`
		TxOffset         int                  `json:"tx_offset"`
		Successful       bool                 `json:"successful"`
		FromAddress      string               `json:"from_address"`
		MinerAddress     string               `json:"miner_address"`
		FromAddressLabel string               `json:"from_address_label"`
		ToAddress        string               `json:"to_address"`
		ToAddressLabel   string               `json:"to_address_label"`
		Value            string               `json:"value"`
		ValueQuote       int                  `json:"value_quote"`
		PrettyValueQuote string               `json:"pretty_value_quote"`
		GasMetadata      GasMetadataPaginated `json:"gas_metadata"`
		GasOffered       int                  `json:"gas_offered"`
		GasSpent         int                  `json:"gas_spent"`
		GasPrice         int                  `json:"gas_price"`
		FeesPaid         string               `json:"fees_paid"`
		GasQuote         int                  `json:"gas_quote"`
		PrettyGasQuote   string               `json:"pretty_gas_quote"`
		GasQuoteRate     int                  `json:"gas_quote_rate"`
		Explorers        []ExplorerPaginated  `json:"explorers"`
		LogEvents        []LogEventPaginated  `json:"log_events"`
	}

	GasMetadataPaginated struct {
		ContractDecimals     int      `json:"contract_decimals"`
		ContractName         string   `json:"contract_name"`
		ContractTickerSymbol string   `json:"contract_ticker_symbol"`
		ContractAddress      string   `json:"contract_address"`
		SupportsERC          []string `json:"supports_erc"`
		LogoURL              string   `json:"logo_url"`
	}

	ExplorerPaginated struct {
		Label string `json:"label"`
		URL   string `json:"url"`
	}

	LogEventPaginated struct {
		BlockSignedAt              string              `json:"block_signed_at"`
		BlockHeight                int                 `json:"block_height"`
		TxOffset                   int                 `json:"tx_offset"`
		LogOffset                  int                 `json:"log_offset"`
		TxHash                     string              `json:"tx_hash"`
		RawLogTopics               []string            `json:"raw_log_topics"`
		SenderContractDecimals     int                 `json:"sender_contract_decimals"`
		SenderName                 string              `json:"sender_name"`
		SenderContractTickerSymbol string              `json:"sender_contract_ticker_symbol"`
		SenderAddress              string              `json:"sender_address"`
		SenderAddressLabel         string              `json:"sender_address_label"`
		SenderLogoURL              string              `json:"sender_logo_url"`
		SupportsERC                []string            `json:"supports_erc"`
		SenderFactoryAddress       string              `json:"sender_factory_address"`
		RawLogData                 string              `json:"raw_log_data"`
		Decoded                    DecodedLogPaginated `json:"decoded"`
	}

	DecodedLogPaginated struct {
		Name      string                  `json:"name"`
		Signature string                  `json:"signature"`
		Params    []DecodedParamPaginated `json:"params"`
	}

	DecodedParamPaginated struct {
		Name    string `json:"name"`
		Type    string `json:"type"`
		Indexed bool   `json:"indexed"`
		Decoded bool   `json:"decoded"`
		Value   string `json:"value"`
	}
)

type (
	ReqGetBulkTimeBucketForAddress struct {
		WalletAddress string `json:"wallet_address"`
		ChainName     string `json:"chain_name"`
		TimeBucket    int    `json:"page"`
	}

	ResGetBulkTimeBucketForAddress struct {
		Address       string        `json:"address"`
		UpdatedAt     time.Time     `json:"updated_at"`
		QuoteCurrency string        `json:"quote_currency"`
		ChainID       int           `json:"chain_id"`
		ChainName     string        `json:"chain_name"`
		Complete      bool          `json:"complete"`
		CurrentBucket int           `json:"current_bucket"`
		Links         LinksBulkTime `json:"links"`
		Items         []TxItem      `json:"items"`
	}

	LinksBulkTime struct {
		Prev string `json:"prev"`
		Next string `json:"next"`
	}

	TxItem struct {
		BlockSignedAt    time.Time           `json:"block_signed_at"`
		BlockHeight      int                 `json:"block_height"`
		BlockHash        string              `json:"block_hash"`
		TxHash           string              `json:"tx_hash"`
		TxOffset         int                 `json:"tx_offset"`
		Successful       bool                `json:"successful"`
		FromAddress      string              `json:"from_address"`
		MinerAddress     string              `json:"miner_address"`
		FromAddressLabel string              `json:"from_address_label"`
		ToAddress        string              `json:"to_address"`
		ToAddressLabel   string              `json:"to_address_label"`
		Value            string              `json:"value"`
		ValueQuote       int                 `json:"value_quote"`
		PrettyValueQuote string              `json:"pretty_value_quote"`
		GasMetadata      GasMetadataBulkTime `json:"gas_metadata"`
		GasOffered       int                 `json:"gas_offered"`
		GasSpent         int                 `json:"gas_spent"`
		GasPrice         int                 `json:"gas_price"`
		FeesPaid         string              `json:"fees_paid"`
		GasQuote         int                 `json:"gas_quote"`
		PrettyGasQuote   string              `json:"pretty_gas_quote"`
		GasQuoteRate     int                 `json:"gas_quote_rate"`
		Explorers        []ExplorerBulkTime  `json:"explorers"`
		LogEvents        []LogEventBulkTime  `json:"log_events"`
	}

	GasMetadataBulkTime struct {
		ContractDecimals     int      `json:"contract_decimals"`
		ContractName         string   `json:"contract_name"`
		ContractTickerSymbol string   `json:"contract_ticker_symbol"`
		ContractAddress      string   `json:"contract_address"`
		SupportsERC          []string `json:"supports_erc"`
		LogoURL              string   `json:"logo_url"`
	}

	ExplorerBulkTime struct {
		Label string `json:"label"`
		URL   string `json:"url"`
	}

	LogEventBulkTime struct {
		BlockSignedAt              time.Time           `json:"block_signed_at"`
		BlockHeight                int                 `json:"block_height"`
		TxOffset                   int                 `json:"tx_offset"`
		LogOffset                  int                 `json:"log_offset"`
		TxHash                     string              `json:"tx_hash"`
		RawLogTopics               []string            `json:"raw_log_topics"`
		SenderContractDecimals     int                 `json:"sender_contract_decimals"`
		SenderName                 string              `json:"sender_name"`
		SenderContractTickerSymbol string              `json:"sender_contract_ticker_symbol"`
		SenderAddress              string              `json:"sender_address"`
		SenderAddressLabel         string              `json:"sender_address_label"`
		SenderLogoURL              string              `json:"sender_logo_url"`
		SupportsERC                []string            `json:"supports_erc"`
		SenderFactoryAddress       string              `json:"sender_factory_address"`
		RawLogData                 string              `json:"raw_log_data"`
		Decoded                    DecodedDataBulkTime `json:"decoded"`
	}

	DecodedDataBulkTime struct {
		Name      string          `json:"name"`
		Signature string          `json:"signature"`
		Params    []ParamBulkTime `json:"params"`
	}

	ParamBulkTime struct {
		Name    string `json:"name"`
		Type    string `json:"type"`
		Indexed bool   `json:"indexed"`
		Decoded bool   `json:"decoded"`
		Value   string `json:"value"`
	}
)

type (
	ReqGetAllTransactionsBlockByPage struct {
		ChainName   string `json:"chainName"`
		BlockHeight int    `json:"block_height"`
		Page        int    `json:"page"`
	}

	ResGetAllTransactionsBlockByPage struct {
		UpdatedAt time.Time           `json:"updated_at"`
		ChainID   int                 `json:"chain_id"`
		ChainName string              `json:"chain_name"`
		Links     LinksBlockByPage    `json:"links"`
		Items     []TxItemBlockByPage `json:"items"`
	}

	LinksBlockByPage struct {
		Prev string `json:"prev"`
		Next string `json:"next"`
	}

	TxItemBlockByPage struct {
		BlockSignedAt    time.Time              `json:"block_signed_at"`
		BlockHeight      int                    `json:"block_height"`
		BlockHash        string                 `json:"block_hash"`
		TxHash           string                 `json:"tx_hash"`
		TxOffset         int                    `json:"tx_offset"`
		Successful       bool                   `json:"successful"`
		FromAddress      string                 `json:"from_address"`
		MinerAddress     string                 `json:"miner_address"`
		FromAddressLabel string                 `json:"from_address_label"`
		ToAddress        string                 `json:"to_address"`
		ToAddressLabel   string                 `json:"to_address_label"`
		Value            string                 `json:"value"`
		ValueQuote       int                    `json:"value_quote"`
		PrettyValueQuote string                 `json:"pretty_value_quote"`
		GasMetadata      GasMetadataBlockByPage `json:"gas_metadata"`
		GasOffered       int                    `json:"gas_offered"`
		GasSpent         int                    `json:"gas_spent"`
		GasPrice         int                    `json:"gas_price"`
		FeesPaid         string                 `json:"fees_paid"`
		GasQuote         int                    `json:"gas_quote"`
		PrettyGasQuote   string                 `json:"pretty_gas_quote"`
		GasQuoteRate     int                    `json:"gas_quote_rate"`
		Explorers        []ExplorerBlockByPage  `json:"explorers"`
		LogEvents        []LogEventBlockByPage  `json:"log_events"`
	}

	GasMetadataBlockByPage struct {
		ContractDecimals     int      `json:"contract_decimals"`
		ContractName         string   `json:"contract_name"`
		ContractTickerSymbol string   `json:"contract_ticker_symbol"`
		ContractAddress      string   `json:"contract_address"`
		SupportsERC          []string `json:"supports_erc"`
		LogoURL              string   `json:"logo_url"`
	}

	ExplorerBlockByPage struct {
		Label string `json:"label"`
		URL   string `json:"url"`
	}

	LogEventBlockByPage struct {
		BlockSignedAt              time.Time              `json:"block_signed_at"`
		BlockHeight                int                    `json:"block_height"`
		TxOffset                   int                    `json:"tx_offset"`
		LogOffset                  int                    `json:"log_offset"`
		TxHash                     string                 `json:"tx_hash"`
		RawLogTopics               []string               `json:"raw_log_topics"`
		SenderContractDecimals     int                    `json:"sender_contract_decimals"`
		SenderName                 string                 `json:"sender_name"`
		SenderContractTickerSymbol string                 `json:"sender_contract_ticker_symbol"`
		SenderAddress              string                 `json:"sender_address"`
		SenderAddressLabel         string                 `json:"sender_address_label"`
		SenderLogoURL              string                 `json:"sender_logo_url"`
		SupportsERC                []string               `json:"supports_erc"`
		SenderFactoryAddress       string                 `json:"sender_factory_address"`
		RawLogData                 string                 `json:"raw_log_data"`
		Decoded                    DecodedDataBlockByPage `json:"decoded"`
	}

	DecodedDataBlockByPage struct {
		Name      string             `json:"name"`
		Signature string             `json:"signature"`
		Params    []ParamBlockByPage `json:"params"`
	}

	ParamBlockByPage struct {
		Name    string `json:"name"`
		Type    string `json:"type"`
		Indexed bool   `json:"indexed"`
		Decoded bool   `json:"decoded"`
		Value   string `json:"value"`
	}
)

type (
	ReqGetAllTransactionsInBlock struct {
		ChainName string `json:"chainName"`
		BlockHash string `json:"block_hash"`
	}

	ResGetAllTransactionsInBlock struct {
		UpdatedAt time.Time       `json:"updated_at"`
		ChainID   int             `json:"chain_id"`
		ChainName string          `json:"chain_name"`
		Items     []TxItemInBlock `json:"items"`
	}

	TxItemInBlock struct {
		BlockSignedAt    time.Time          `json:"block_signed_at"`
		BlockHeight      int                `json:"block_height"`
		BlockHash        string             `json:"block_hash"`
		TxHash           string             `json:"tx_hash"`
		TxOffset         int                `json:"tx_offset"`
		Successful       bool               `json:"successful"`
		FromAddress      string             `json:"from_address"`
		MinerAddress     string             `json:"miner_address"`
		FromAddressLabel string             `json:"from_address_label"`
		ToAddress        string             `json:"to_address"`
		ToAddressLabel   string             `json:"to_address_label"`
		Value            string             `json:"value"`
		ValueQuote       int                `json:"value_quote"`
		PrettyValueQuote string             `json:"pretty_value_quote"`
		GasMetadata      GasMetadataInBlock `json:"gas_metadata"`
		GasOffered       int                `json:"gas_offered"`
		GasSpent         int                `json:"gas_spent"`
		GasPrice         int                `json:"gas_price"`
		FeesPaid         string             `json:"fees_paid"`
		GasQuote         int                `json:"gas_quote"`
		PrettyGasQuote   string             `json:"pretty_gas_quote"`
		GasQuoteRate     int                `json:"gas_quote_rate"`
		Explorers        []ExplorerInBlock  `json:"explorers"`
		LogEvents        []LogEventInBlock  `json:"log_events"`
	}

	GasMetadataInBlock struct {
		ContractDecimals     int      `json:"contract_decimals"`
		ContractName         string   `json:"contract_name"`
		ContractTickerSymbol string   `json:"contract_ticker_symbol"`
		ContractAddress      string   `json:"contract_address"`
		SupportsERC          []string `json:"supports_erc"`
		LogoURL              string   `json:"logo_url"`
	}

	ExplorerInBlock struct {
		Label string `json:"label"`
		URL   string `json:"url"`
	}

	LogEventInBlock struct {
		BlockSignedAt              time.Time          `json:"block_signed_at"`
		BlockHeight                int                `json:"block_height"`
		TxOffset                   int                `json:"tx_offset"`
		LogOffset                  int                `json:"log_offset"`
		TxHash                     string             `json:"tx_hash"`
		RawLogTopics               []string           `json:"raw_log_topics"`
		SenderContractDecimals     int                `json:"sender_contract_decimals"`
		SenderName                 string             `json:"sender_name"`
		SenderContractTickerSymbol string             `json:"sender_contract_ticker_symbol"`
		SenderAddress              string             `json:"sender_address"`
		SenderAddressLabel         string             `json:"sender_address_label"`
		SenderLogoURL              string             `json:"sender_logo_url"`
		SupportsERC                []string           `json:"supports_erc"`
		SenderFactoryAddress       string             `json:"sender_factory_address"`
		RawLogData                 string             `json:"raw_log_data"`
		Decoded                    DecodedDataInBlock `json:"decoded"`
	}

	DecodedDataInBlock struct {
		Name      string         `json:"name"`
		Signature string         `json:"signature"`
		Params    []ParamInBlock `json:"params"`
	}

	ParamInBlock struct {
		Name    string `json:"name"`
		Type    string `json:"type"`
		Indexed bool   `json:"indexed"`
		Decoded bool   `json:"decoded"`
		Value   string `json:"value"`
	}
)

type (
	ReqGetTokenApprovalsForAddress struct {
		ChainName     string `json:"chain_name"`
		WalletAddress string `json:"wallet_address"`
	}

	ResGetTokenApprovalsForAddress struct {
		Address       string          `json:"address"`
		UpdatedAt     time.Time       `json:"updated_at"`
		QuoteCurrency string          `json:"quote_currency"`
		ChainID       int             `json:"chain_id"`
		ChainName     string          `json:"chain_name"`
		Items         []TokenApproval `json:"items"`
	}

	TokenApproval struct {
		TokenAddress           string    `json:"token_address"`
		TokenAddressLabel      string    `json:"token_address_label"`
		TickerSymbol           string    `json:"ticker_symbol"`
		ContractDecimals       int       `json:"contract_decimals"`
		LogoURL                string    `json:"logo_url"`
		QuoteRate              int       `json:"quote_rate"`
		Balance                string    `json:"balance"`
		BalanceQuote           int       `json:"balance_quote"`
		PrettyBalanceQuote     string    `json:"pretty_balance_quote"`
		ValueAtRisk            string    `json:"value_at_risk"`
		ValueAtRiskQuote       int       `json:"value_at_risk_quote"`
		PrettyValueAtRiskQuote string    `json:"pretty_value_at_risk_quote"`
		Spenders               []Spender `json:"spenders"`
	}

	Spender struct {
		BlockHeight            int       `json:"block_height"`
		TxOffset               int       `json:"tx_offset"`
		LogOffset              int       `json:"log_offset"`
		BlockSignedAt          time.Time `json:"block_signed_at"`
		TxHash                 string    `json:"tx_hash"`
		SpenderAddress         string    `json:"spender_address"`
		SpenderAddressLabel    string    `json:"spender_address_label"`
		Allowance              string    `json:"allowance"`
		AllowanceQuote         int       `json:"allowance_quote"`
		PrettyAllowanceQuote   string    `json:"pretty_allowance_quote"`
		ValueAtRisk            string    `json:"value_at_risk"`
		ValueAtRiskQuote       int       `json:"value_at_risk_quote"`
		PrettyValueAtRiskQuote string    `json:"pretty_value_at_risk_quote"`
		RiskFactor             string    `json:"risk_factor"`
	}
)

type (
	ReqGetNFTApprovalsForAddress struct {
		ChainName     string `json:"chain_name"`
		WalletAddress string `json:"wallet_address"`
	}

	ResGetNFTApprovalsForAddress struct {
		UpdatedAt time.Time     `json:"updated_at"`
		ChainID   int           `json:"chain_id"`
		ChainName string        `json:"chain_name"`
		Address   string        `json:"address"`
		Items     []NFTApproval `json:"items"`
	}

	NFTApproval struct {
		ContractAddress      string         `json:"contract_address"`
		ContractAddressLabel string         `json:"contract_address_label"`
		ContractTickerSymbol string         `json:"contract_ticker_symbol"`
		TokenBalances        []TokenBalance `json:"token_balances"`
		Spenders             []SpenderNFT   `json:"spenders"`
	}

	TokenBalance struct {
		TokenID      string `json:"token_id"`
		TokenBalance string `json:"token_balance"`
	}

	SpenderNFT struct {
		BlockHeight         int       `json:"block_height"`
		TxOffset            int       `json:"tx_offset"`
		LogOffset           int       `json:"log_offset"`
		BlockSignedAt       time.Time `json:"block_signed_at"`
		TxHash              string    `json:"tx_hash"`
		SpenderAddress      string    `json:"spender_address"`
		SpenderAddressLabel string    `json:"spender_address_label"`
		TokenIDsApproved    string    `json:"token_ids_approved"`
		Allowance           string    `json:"allowance"`
	}
)

type (
	ReqGetBitcoinBalancesForHDAddress struct {
		WalletAddress string `json:"wallet_address"`
	}

	ResGetBitcoinBalancesForHDAddress struct {
		Address       string               `json:"address"`
		ChainID       int                  `json:"chain_id"`
		ChainName     string               `json:"chain_name"`
		QuoteCurrency string               `json:"quote_currency"`
		UpdatedAt     time.Time            `json:"updated_at"`
		Items         []BitcoinBalanceItem `json:"items"`
	}

	BitcoinBalanceItem struct {
		ChildAddress         string                       `json:"child_address"`
		AddressPath          string                       `json:"address_path"`
		ContractDecimals     int                          `json:"contract_decimals"`
		ContractName         string                       `json:"contract_name"`
		ContractTickerSymbol string                       `json:"contract_ticker_symbol"`
		ContractAddress      string                       `json:"contract_address"`
		ContractDisplayName  string                       `json:"contract_display_name"`
		SupportsERC          []string                     `json:"supports_erc"`
		LogoURL              string                       `json:"logo_url"`
		LogoURLs             LogoURLsBalancesForHDAddress `json:"logo_urls"`
		LastTransferredAt    time.Time                    `json:"last_transferred_at"`
		NativeToken          bool                         `json:"native_token"`
		Type                 string                       `json:"type"`
		IsSpam               bool                         `json:"is_spam"`
		Balance              string                       `json:"balance"`
		Balance24h           string                       `json:"balance_24h"`
		QuoteRate            float64                      `json:"quote_rate"`
		QuoteRate24h         float64                      `json:"quote_rate_24h"`
		Quote                float64                      `json:"quote"`
		Quote24h             float64                      `json:"quote_24h"`
		PrettyQuote          string                       `json:"pretty_quote"`
		PrettyQuote24h       string                       `json:"pretty_quote_24h"`
		ProtocolMetadata     ProtocolData                 `json:"protocol_metadata"`
		NFTData              string                       `json:"nft_data"`
	}

	LogoURLsBalancesForHDAddress struct {
		TokenLogoURL    string `json:"token_logo_url"`
		ProtocolLogoURL string `json:"protocol_logo_url"`
		ChainLogoURL    string `json:"chain_logo_url"`
	}
	ProtocolData struct {
		ProtocolName string `json:"protocol_name"`
	}
)

type (
	ReqGetBitcoinBalancesForNonHDAddress struct {
		WalletAddress string `json:"wallet_address"`
	}

	ResGetBitcoinBalancesForNonHDAddress struct {
		Address       string                    `json:"address"`
		ChainID       int                       `json:"chain_id"`
		ChainName     string                    `json:"chain_name"`
		QuoteCurrency string                    `json:"quote_currency"`
		UpdatedAt     time.Time                 `json:"updated_at"`
		Items         []BitcoinNonHDBalanceItem `json:"items"`
	}

	BitcoinNonHDBalanceItem struct {
		ContractDecimals     int                      `json:"contract_decimals"`
		ContractName         string                   `json:"contract_name"`
		ContractTickerSymbol string                   `json:"contract_ticker_symbol"`
		ContractAddress      string                   `json:"contract_address"`
		ContractDisplayName  string                   `json:"contract_display_name"`
		SupportsERC          []string                 `json:"supports_erc"`
		LogoURL              string                   `json:"logo_url"`
		LogoURLs             LogoURLsNonHDAddress     `json:"logo_urls"`
		LastTransferredAt    time.Time                `json:"last_transferred_at"`
		NativeToken          bool                     `json:"native_token"`
		Type                 string                   `json:"type"`
		IsSpam               bool                     `json:"is_spam"`
		Balance              string                   `json:"balance"`
		Balance24h           string                   `json:"balance_24h"`
		QuoteRate            float64                  `json:"quote_rate"`
		QuoteRate24h         float64                  `json:"quote_rate_24h"`
		Quote                float64                  `json:"quote"`
		Quote24h             float64                  `json:"quote_24h"`
		PrettyQuote          string                   `json:"pretty_quote"`
		PrettyQuote24h       string                   `json:"pretty_quote_24h"`
		ProtocolMetadata     ProtocolDataNonHDAddress `json:"protocol_metadata"`
		NFTData              []NFTDataNonHDAddress    `json:"nft_data"`
	}

	LogoURLsNonHDAddress struct {
		TokenLogoURL    string `json:"token_logo_url"`
		ProtocolLogoURL string `json:"protocol_logo_url"`
		ChainLogoURL    string `json:"chain_logo_url"`
	}

	ProtocolDataNonHDAddress struct {
		ProtocolName string `json:"protocol_name"`
	}

	NFTDataNonHDAddress struct {
		TokenID           string                   `json:"token_id"`
		TokenBalance      string                   `json:"token_balance"`
		TokenURL          string                   `json:"token_url"`
		SupportsERC       []string                 `json:"supports_erc"`
		TokenPriceWei     string                   `json:"token_price_wei"`
		TokenQuoteRateETH string                   `json:"token_quote_rate_eth"`
		OriginalOwner     string                   `json:"original_owner"`
		ExternalData      ExternalDataNonHDAddress `json:"external_data"`
		Owner             string                   `json:"owner"`
		OwnerAddress      string                   `json:"owner_address"`
		Burned            bool                     `json:"burned"`
	}

	ExternalDataNonHDAddress struct {
		Name            string      `json:"name"`
		Description     string      `json:"description"`
		Image           string      `json:"image"`
		Image256        string      `json:"image_256"`
		Image512        string      `json:"image_512"`
		Image1024       string      `json:"image_1024"`
		AnimationURL    string      `json:"animation_url"`
		ExternalURL     string      `json:"external_url"`
		Attributes      []Attribute `json:"attributes"`
		Thumbnails      Thumbnails  `json:"thumbnails"`
		ImagePreview    string      `json:"image_preview"`
		AssetProperties AssetProps  `json:"asset_properties"`
		Owner           string      `json:"owner"`
	}

	Attribute struct {
		TraitType string `json:"trait_type"`
		Value     string `json:"value"`
	}

	Thumbnails struct {
		Image256       string `json:"image_256"`
		Image512       string `json:"image_512"`
		Image1024      string `json:"image_1024"`
		ImageOpenGraph string `json:"image_opengraph_url"`
		Thumbhash      string `json:"thumbhash"`
	}

	AssetProps struct {
		AssetWidth    int    `json:"asset_width"`
		AssetHeight   int    `json:"asset_height"`
		DominantColor string `json:"dominant_color"`
	}
)

type (
	ReqGetTransactionsForBitcoinAddress struct {
	}

	ResGetTransactionsForBitcoinAddress struct {
		UpdatedAt  time.Time                   `json:"updated_at"`
		Items      []BitcoinTransactionItem    `json:"items"`
		Pagination PaginationForBitcoinAddress `json:"pagination"`
	}

	BitcoinTransactionItem struct {
		ChainID          int       `json:"chain_id"`
		ChainName        string    `json:"chain_name"`
		ContractDecimals int       `json:"contract_decimals"`
		BlockSignedAt    time.Time `json:"block_signed_at"`
		BlockHeight      int       `json:"block_height"`
		BlockHash        string    `json:"block_hash"`
		TxHash           string    `json:"tx_hash"`
		TxIdx            int       `json:"tx_idx"`
		Type             string    `json:"type"`
		Address          string    `json:"address"`
		Value            string    `json:"value"`
		Quote            float64   `json:"quote"`
		QuoteRate        float64   `json:"quote_rate"`
		FeesPaid         string    `json:"fees_paid"`
		GasQuote         float64   `json:"gas_quote"`
		GasQuoteRate     float64   `json:"gas_quote_rate"`
		Coinbase         bool      `json:"coinbase"`
		Locktime         int       `json:"locktime"`
		Weight           int       `json:"weight"`
	}

	PaginationForBitcoinAddress struct {
		HasMore    bool `json:"has_more"`
		PageNumber int  `json:"page_number"`
		PageSize   int  `json:"page_size"`
		TotalCount int  `json:"total_count"`
	}
)

type (
	ReqGetHistoricalTokenPrice struct {
		ChainName       string `json:"chainName"`
		QuoteCurrency   string `json:"quoteCurrency"`
		ContractAddress string `json:"contractAddress"`
	}

	ResGetHistoricalTokenPrice struct {
		ContractDecimals     int                          `json:"contract_decimals"`
		ContractName         string                       `json:"contract_name"`
		ContractTickerSymbol string                       `json:"contract_ticker_symbol"`
		ContractAddress      string                       `json:"contract_address"`
		SupportsERC          []string                     `json:"supports_erc"`
		LogoURL              string                       `json:"logo_url"`
		UpdatedAt            time.Time                    `json:"update_at"`
		QuoteCurrency        string                       `json:"quote_currency"`
		LogoURLs             LogoURLsHistoricalTokenPrice `json:"logo_urls"`
		Prices               []PriceItem                  `json:"prices"`
		Items                []interface{}                `json:"items"`
	}

	LogoURLsHistoricalTokenPrice struct {
		TokenLogoURL    string `json:"token_logo_url"`
		ProtocolLogoURL string `json:"protocol_logo_url"`
		ChainLogoURL    string `json:"chain_logo_url"`
	}

	PriceItem struct {
		ContractMetadata ContractMetadata `json:"contract_metadata"`
		Date             time.Time        `json:"date"`
		Price            float64          `json:"price"`
		PrettyPrice      string           `json:"pretty_price"`
	}

	ContractMetadata struct {
		ContractDecimals     int      `json:"contract_decimals"`
		ContractName         string   `json:"contract_name"`
		ContractTickerSymbol string   `json:"contract_ticker_symbol"`
		ContractAddress      string   `json:"contract_address"`
		SupportsERC          []string `json:"supports_erc"`
		LogoURL              string   `json:"logo_url"`
	}
)

type (
	ReqGetPoolSpotPrices struct {
		ChainName       string `json:"chainName"`
		ContractAddress string `json:"contractAddress"`
	}

	ResGetPoolSpotPrices struct {
		UpdatedAt           time.Time `json:"updated_at"`
		PoolAddress         string    `json:"pool_address"`
		Token0Address       string    `json:"token_0_address"`
		Token0Name          string    `json:"token_0_name"`
		Token0Ticker        string    `json:"token_0_ticker"`
		Token0Price         string    `json:"token_0_price"`
		Token0Price24h      string    `json:"token_0_price_24h"`
		Token0PriceQuote    string    `json:"token_0_price_quote"`
		Token0Price24hQuote string    `json:"token_0_price_24h_quote"`
		Token1Address       string    `json:"token_1_address"`
		Token1Name          string    `json:"token_1_name"`
		Token1Ticker        string    `json:"token_1_ticker"`
		Token1Price         string    `json:"token_1_price"`
		Token1Price24h      string    `json:"token_1_price_24h"`
		Token1PriceQuote    string    `json:"token_1_price_quote"`
		Token1Price24hQuote string    `json:"token_1_price_24h_quote"`
		QuoteCurrency       string    `json:"quote_currency"`
	}
)

type (
	ReqGetGasPrices struct {
		ChainName string `json:"chainName"`
		EventType string `json:"eventType"`
	}

	ResGetGasPrices struct {
		ChainID       int       `json:"chain_id"`
		ChainName     string    `json:"chain_name"`
		QuoteCurrency string    `json:"quote_currency"`
		UpdatedAt     time.Time `json:"updated_at"`
		EventType     string    `json:"event_type"`
		GasQuoteRate  int       `json:"gas_quote_rate"`
		BaseFee       string    `json:"base_fee"`
		Items         []GasItem `json:"items"`
	}

	GasItem struct {
		GasPrice            string    `json:"gas_price"`
		GasSpent            string    `json:"gas_spent"`
		GasQuote            int       `json:"gas_quote"`
		OtherFees           OtherFees `json:"other_fees"`
		TotalGasQuote       int       `json:"total_gas_quote"`
		PrettyTotalGasQuote string    `json:"pretty_total_gas_quote"`
		Interval            string    `json:"interval"`
	}

	OtherFees struct {
		L1GasQuote int `json:"l_1_gas_quote"`
	}
)

type (
	ReqGetLogs struct {
		ChainName string `json:"chainName"`
	}

	ResGetLogs struct {
		UpdatedAt time.Time `json:"updated_at"`
		ChainID   int       `json:"chain_id"`
		ChainName string    `json:"chain_name"`
		Items     []LogItem `json:"items"`
	}

	LogItem struct {
		BlockSignedAt              time.Time         `json:"block_signed_at"`
		BlockHeight                int               `json:"block_height"`
		BlockHash                  string            `json:"block_hash"`
		TxOffset                   int               `json:"tx_offset"`
		LogOffset                  int               `json:"log_offset"`
		TxHash                     string            `json:"tx_hash"`
		RawLogTopics               []string          `json:"raw_log_topics"`
		SenderContractDecimals     int               `json:"sender_contract_decimals"`
		SenderName                 string            `json:"sender_name"`
		SenderContractTickerSymbol string            `json:"sender_contract_ticker_symbol"`
		SenderAddress              string            `json:"sender_address"`
		SenderAddressLabel         string            `json:"sender_address_label"`
		SupportsERC                []string          `json:"supports_erc"`
		SenderLogoURL              string            `json:"sender_logo_url"`
		SenderFactoryAddress       string            `json:"sender_factory_address"`
		RawLogData                 string            `json:"raw_log_data"`
		Decoded                    DecodedLogGetLogs `json:"decoded"`
	}

	DecodedLogGetLogs struct {
		Name      string                `json:"name"`
		Signature string                `json:"signature"`
		Params    []DecodedParamGetLogs `json:"params"`
	}

	DecodedParamGetLogs struct {
		Name    string `json:"name"`
		Type    string `json:"type"`
		Indexed bool   `json:"indexed"`
		Decoded bool   `json:"decoded"`
		Value   string `json:"value"`
	}
)

type (
	ReqGetLogEventsByTopicHash struct {
		ChainName string `json:"chainName"`
		TopicHash string `json:"topic_hash"`
	}

	ResGetLogEventsByTopicHash struct {
		UpdatedAt  time.Time             `json:"updated_at"`
		ChainID    int                   `json:"chain_id"`
		ChainName  string                `json:"chain_name"`
		Items      []LogEventByTopicHash `json:"items"`
		Pagination PaginationByTopicHash `json:"pagination"`
	}

	LogEventByTopicHash struct {
		BlockSignedAt              time.Time             `json:"block_signed_at"`
		BlockHeight                int                   `json:"block_height"`
		TxOffset                   int                   `json:"tx_offset"`
		LogOffset                  int                   `json:"log_offset"`
		TxHash                     string                `json:"tx_hash"`
		RawLogTopics               []string              `json:"raw_log_topics"`
		SenderContractDecimals     int                   `json:"sender_contract_decimals"`
		SenderName                 string                `json:"sender_name"`
		SenderContractTickerSymbol string                `json:"sender_contract_ticker_symbol"`
		SenderAddress              string                `json:"sender_address"`
		SenderAddressLabel         string                `json:"sender_address_label"`
		SenderLogoURL              string                `json:"sender_logo_url"`
		SupportsERC                []string              `json:"supports_erc"`
		SenderFactoryAddress       string                `json:"sender_factory_address"`
		RawLogData                 string                `json:"raw_log_data"`
		Decoded                    DecodedLogByTopicHash `json:"decoded"`
	}

	DecodedLogByTopicHash struct {
		Name      string                    `json:"name"`
		Signature string                    `json:"signature"`
		Params    []DecodedParamByTopicHash `json:"params"`
	}

	DecodedParamByTopicHash struct {
		Name    string `json:"name"`
		Type    string `json:"type"`
		Indexed bool   `json:"indexed"`
		Decoded bool   `json:"decoded"`
		Value   string `json:"value"`
	}

	PaginationByTopicHash struct {
		HasMore    bool `json:"has_more"`
		PageNumber int  `json:"page_number"`
		PageSize   int  `json:"page_size"`
		TotalCount int  `json:"total_count"`
	}
)

type (
	ReqGetLogEventsByContractAddress struct {
		ChainName       string `json:"chainName"`
		ContractAddress string `json:"contract_address"`
	}

	ResGetLogEventsByContractAddress struct {
		UpdatedAt  time.Time                   `json:"updated_at"`
		ChainID    int                         `json:"chain_id"`
		ChainName  string                      `json:"chain_name"`
		Items      []LogEventByContractAddress `json:"items"`
		Pagination PaginationByContractAddress `json:"pagination"`
	}

	LogEventByContractAddress struct {
		BlockSignedAt              time.Time                   `json:"block_signed_at"`
		BlockHeight                int                         `json:"block_height"`
		TxOffset                   int                         `json:"tx_offset"`
		LogOffset                  int                         `json:"log_offset"`
		TxHash                     string                      `json:"tx_hash"`
		RawLogTopics               []string                    `json:"raw_log_topics"`
		SenderContractDecimals     int                         `json:"sender_contract_decimals"`
		SenderName                 string                      `json:"sender_name"`
		SenderContractTickerSymbol string                      `json:"sender_contract_ticker_symbol"`
		SenderAddress              string                      `json:"sender_address"`
		SenderAddressLabel         string                      `json:"sender_address_label"`
		SenderLogoURL              string                      `json:"sender_logo_url"`
		SupportsERC                []string                    `json:"supports_erc"`
		SenderFactoryAddress       string                      `json:"sender_factory_address"`
		RawLogData                 string                      `json:"raw_log_data"`
		Decoded                    DecodedLogByContractAddress `json:"decoded"`
	}

	DecodedLogByContractAddress struct {
		Name      string                          `json:"name"`
		Signature string                          `json:"signature"`
		Params    []DecodedParamByContractAddress `json:"params"`
	}

	DecodedParamByContractAddress struct {
		Name    string `json:"name"`
		Type    string `json:"type"`
		Indexed bool   `json:"indexed"`
		Decoded bool   `json:"decoded"`
		Value   string `json:"value"`
	}

	PaginationByContractAddress struct {
		HasMore    bool `json:"has_more"`
		PageNumber int  `json:"page_number"`
		PageSize   int  `json:"page_size"`
		TotalCount int  `json:"total_count"`
	}
)

type (
	ReqGetABlock struct {
		ChainName   string `json:"chain_name"`
		BlockHeight string `json:"block_height"`
	}

	ResGetABlock struct {
		UpdatedAt time.Time `json:"updated_at"`
		ChainID   int       `json:"chain_id"`
		ChainName string    `json:"chain_name"`
		Items     []Block   `json:"items"`
	}

	Block struct {
		BlockHash        string    `json:"block_hash"`
		SignedAt         time.Time `json:"signed_at"`
		Height           int       `json:"height"`
		BlockParentHash  string    `json:"block_parent_hash"`
		ExtraData        string    `json:"extra_data"`
		MinerAddress     string    `json:"miner_address"`
		MiningCost       int       `json:"mining_cost"`
		GasUsed          int       `json:"gas_used"`
		GasLimit         int       `json:"gas_limit"`
		TransactionsLink string    `json:"transactions_link"`
	}
)

type (
	ReqGetBlockHeights struct {
		ChainName string `json:"chain_name"`
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
	}

	ResGetBlockHeights struct {
		UpdatedAt  time.Time      `json:"updated_at"`
		ChainID    int            `json:"chain_id"`
		ChainName  string         `json:"chain_name"`
		Items      []BlockHeight  `json:"items"`
		Pagination PaginationInfo `json:"pagination"`
	}

	BlockHeight struct {
		BlockHash        string    `json:"block_hash"`
		SignedAt         time.Time `json:"signed_at"`
		Height           int       `json:"height"`
		BlockParentHash  string    `json:"block_parent_hash"`
		ExtraData        string    `json:"extra_data"`
		MinerAddress     string    `json:"miner_address"`
		MiningCost       int       `json:"mining_cost"`
		GasUsed          int       `json:"gas_used"`
		GasLimit         int       `json:"gas_limit"`
		TransactionsLink string    `json:"transactions_link"`
	}

	PaginationInfo struct {
		HasMore    bool `json:"has_more"`
		PageNumber int  `json:"page_number"`
		PageSize   int  `json:"page_size"`
		TotalCount int  `json:"total_count"`
	}
)

type (
	ReqGetAllChains struct {
	}

	ResGetAllChains struct {
		UpdatedAt time.Time   `json:"updated_at"`
		Items     []ChainInfo `json:"items"`
	}

	ChainInfo struct {
		Name          string      `json:"name"`
		ChainID       string      `json:"chain_id"`
		IsTestnet     bool        `json:"is_testnet"`
		DBSchemaName  string      `json:"db_schema_name"`
		Label         string      `json:"label"`
		CategoryLabel string      `json:"category_label"`
		LogoURL       string      `json:"logo_url"`
		BlackLogoURL  string      `json:"black_logo_url"`
		WhiteLogoURL  string      `json:"white_logo_url"`
		ColorTheme    ColorTheme  `json:"color_theme"`
		IsAppchain    bool        `json:"is_appchain"`
		AppchainOf    interface{} `json:"appchain_of"`
	}

	ColorTheme struct {
		Red    int    `json:"red"`
		Green  int    `json:"green"`
		Blue   int    `json:"blue"`
		Alpha  int    `json:"alpha"`
		Hex    string `json:"hex"`
		CSSRGB string `json:"css_rgb"`
	}
)

type (
	ReqGetAllChainStatuses struct {
	}

	ResGetAllChainStatuses struct {
		UpdatedAt time.Time     `json:"updated_at"`
		Items     []ChainStatus `json:"items"`
	}

	ChainStatus struct {
		Name                  string    `json:"name"`
		ChainID               string    `json:"chain_id"`
		IsTestnet             bool      `json:"is_testnet"`
		LogoURL               string    `json:"logo_url"`
		BlackLogoURL          string    `json:"black_logo_url"`
		WhiteLogoURL          string    `json:"white_logo_url"`
		IsAppchain            bool      `json:"is_appchain"`
		SyncedBlockHeight     int       `json:"synced_block_height"`
		SyncedBlockedSignedAt time.Time `json:"synced_blocked_signed_at"`
		HasData               bool      `json:"has_data"`
	}
)

type (
	ReqGetResolverAddressForRegisteredAddress struct {
		ChainName     string `json:"chain_name"`
		WalletAddress string `json:"wallet_address"`
	}

	ResGetResolverAddressForRegisteredAddress struct {
		UpdatedAt time.Time             `json:"updated_at"`
		ChainID   int                   `json:"chain_id"`
		ChainName string                `json:"chain_name"`
		Items     []ResolverAddressItem `json:"items"`
	}

	ResolverAddressItem struct {
		Address string `json:"address"`
		Name    string `json:"name"`
	}
)
