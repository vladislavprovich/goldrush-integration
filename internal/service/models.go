package service

type (
	RegisterUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	RegisterUserResponse struct {
		UserID int64 `json:"user_id"`
	}
)

type (
	LoginUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginUserResponse struct {
		Token string `json:"token"`
	}
)

type (
	RecentAddressTransactionRequest struct {
		WalletAddress string `json:"wallet_address"`
		ChainName     string `json:"chain_name"`
	}

	RecentAddressTransactionResponse struct {
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
		ValueQuote       float64           `json:"value_quote"`
		PrettyValueQuote string            `json:"pretty_value_quote"`
		GasMetadata      GasMetadataRecent `json:"gas_metadata"`
		GasOffered       int               `json:"gas_offered"`
		GasSpent         int               `json:"gas_spent"`
		GasPrice         int               `json:"gas_price"`
		FeesPaid         string            `json:"fees_paid"`
		GasQuote         float64           `json:"gas_quote"`
		PrettyGasQuote   string            `json:"pretty_gas_quote"`
		GasQuoteRate     float64           `json:"gas_quote_rate"`
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
		Name    string      `json:"name"`
		Type    string      `json:"type"`
		Indexed bool        `json:"indexed"`
		Decoded bool        `json:"decoded"`
		Value   interface{} `json:"value"`
	}
)

type (
	TransactionInfoRequest struct {
		ChainName string `json:"chain_name"`
		TxHash    string `json:"tx_hash"`
	}

	TransactionInfoResponse struct {
		BlockSignedAt     string                 `json:"block_signed_at"`
		BlockHeight       int64                  `json:"block_height"`
		BlockHash         string                 `json:"block_hash"`
		TxHash            string                 `json:"tx_hash"`
		TxOffset          int64                  `json:"tx_offset"`
		Successful        bool                   `json:"successful"`
		FromAddress       string                 `json:"from_address"`
		MinerAddress      string                 `json:"miner_address"`
		FromAddressLabel  string                 `json:"from_address_label"`
		ToAddress         string                 `json:"to_address"`
		ToAddressLabel    string                 `json:"to_address_label"`
		Value             string                 `json:"value"`
		ValueQuote        float64                `json:"value_quote"`
		PrettyValueQuote  string                 `json:"pretty_value_quote"`
		GasMetadata       GasMetadataInfo        `json:"gas_metadata"`
		GasOffered        int64                  `json:"gas_offered"`
		GasSpent          int64                  `json:"gas_spent"`
		GasPrice          int64                  `json:"gas_price"`
		FeesPaid          string                 `json:"fees_paid"`
		GasQuote          float64                `json:"gas_quote"`
		PrettyGasQuote    string                 `json:"pretty_gas_quote"`
		GasQuoteRate      float64                `json:"gas_quote_rate"`
		ChainID           int64                  `json:"chain_id"`
		ChainName         string                 `json:"chain_name"`
		Explorers         []ExplorerInfo         `json:"explorers"`
		LogEvents         []LogEventInfo         `json:"log_events"`
		InternalTransfers []InternalTransferInfo `json:"internal_transfers"`
		StateChanges      []StateChangeInfo      `json:"state_changes"`
		InputData         InputDataInfo          `json:"input_data"`
	}

	GasMetadataInfo struct {
		ContractDecimals     int64    `json:"contract_decimals"`
		ContractName         string   `json:"contract_name"`
		ContractTickerSymbol string   `json:"contract_ticker_symbol"`
		ContractAddress      string   `json:"contract_address"`
		SupportsERC          []string `json:"supports_erc"`
		LogoURL              string   `json:"logo_url"`
	}

	ExplorerInfo struct {
		Label string `json:"label"`
		URL   string `json:"url"`
	}

	LogEventInfo struct {
		BlockSignedAt          string          `json:"block_signed_at"`
		BlockHeight            int64           `json:"block_height"`
		TxOffset               int64           `json:"tx_offset"`
		LogOffset              int64           `json:"log_offset"`
		TxHash                 string          `json:"tx_hash"`
		RawLogTopics           []string        `json:"raw_log_topics"`
		SenderContractDecimals int64           `json:"sender_contract_decimals"`
		SenderName             string          `json:"sender_name"`
		SenderTickerSymbol     string          `json:"sender_ticker_symbol"`
		SenderAddress          string          `json:"sender_address"`
		SenderAddressLabel     string          `json:"sender_address_label"`
		SenderLogoURL          string          `json:"sender_logo_url"`
		SupportsERC            []string        `json:"supports_erc"`
		SenderFactoryAddress   string          `json:"sender_factory_address"`
		RawLogData             string          `json:"raw_log_data"`
		Decoded                DecodedDataInfo `json:"decoded"`
	}

	DecodedDataInfo struct {
		Name      string      `json:"name"`
		Signature string      `json:"signature"`
		Params    []ParamInfo `json:"params"`
	}

	ParamInfo struct {
		Name    string      `json:"name"`
		Type    string      `json:"type"`
		Indexed bool        `json:"indexed"`
		Decoded bool        `json:"decoded"`
		Value   interface{} `json:"value"`
	}

	InternalTransferInfo struct {
		FromAddress string `json:"from_address"`
		ToAddress   string `json:"to_address"`
		Value       string `json:"value"`
		GasLimit    int64  `json:"gas_limit"`
	}

	StorageChangeInfo struct {
		StorageAddress string `json:"storage_address"`
		ValueBefore    string `json:"value_before"`
		ValueAfter     string `json:"value_after"`
	}

	StateChangeInfo struct {
		Address        string              `json:"address"`
		BalanceBefore  string              `json:"balance_before"`
		BalanceAfter   string              `json:"balance_after"`
		StorageChanges []StorageChangeInfo `json:"storage_changes"`
		NonceBefore    int64               `json:"nonce_before"`
		NonceAfter     int64               `json:"nonce_after"`
	}

	InputDataInfo struct {
		MethodID string `json:"method_id"`
	}
)

type (
	HistoricalPortfolioValueRequest struct {
		WalletAddress string `json:"wallet_address"`
		ChainName     string `json:"chain_name"`
	}

	HistoricalPortfolioValueResponse struct {
		Address       string                    `json:"address"`
		UpdatedAt     string                    `json:"updated_at"`
		QuoteCurrency string                    `json:"quote_currency"`
		ChainID       int                       `json:"chain_id"`
		ChainName     string                    `json:"chain_name"`
		Items         []HistoricalPortfolioItem `json:"items"`
	}

	HistoricalPortfolioItem struct {
		ContractAddress      string              `json:"contract_address"`
		ContractDecimals     int                 `json:"contract_decimals"`
		ContractName         string              `json:"contract_name"`
		ContractTickerSymbol string              `json:"contract_ticker_symbol"`
		LogoURL              string              `json:"logo_url"`
		Holdings             []HistoricalHolding `json:"holdings"`
	}

	HistoricalHolding struct {
		QuoteRate float64           `json:"quote_rate"`
		Timestamp string            `json:"timestamp"`
		Close     HistoricalBalance `json:"close"`
		High      HistoricalBalance `json:"high"`
		Low       HistoricalBalance `json:"low"`
		Open      HistoricalBalance `json:"open"`
	}

	HistoricalBalance struct {
		Balance     string  `json:"balance"`
		Quote       float64 `json:"quote"`
		PrettyQuote string  `json:"pretty_quote"`
	}
)

type (
	TokenBalancesRequest struct {
		WalletAddress string `json:"wallet_address"`
		ChainName     string `json:"chain_name"`
	}

	TokenBalancesResponse struct {
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
		QuoteRate            float64          `json:"quote_rate"`
		QuoteRate24h         float64          `json:"quote_rate_24h"`
		Quote                float64          `json:"quote"`
		Quote24h             float64          `json:"quote_24h"`
		PrettyQuote          string           `json:"pretty_quote"`
		PrettyQuote24h       string           `json:"pretty_quote_24h"`
		ProtocolMetadata     ProtocolMetadata `json:"protocol_metadata"`
		NFTData              []NFTData        `json:"nft_data"`
	}
)

type (
	PricePoint struct {
		Balance     string  `json:"balance"`
		Quote       float64 `json:"quote"`
		PrettyQuote string  `json:"pretty_quote"`
	}

	Holding struct {
		QuoteRate float64    `json:"quote_rate"`
		Timestamp string     `json:"timestamp"`
		Close     PricePoint `json:"close"`
		High      PricePoint `json:"high"`
		Low       PricePoint `json:"low"`
		Open      PricePoint `json:"open"`
	}
)
