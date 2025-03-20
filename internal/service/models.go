package service

type (
	RegisterUserRequest struct {
		Email    string
		Password string
	}

	RegisterUserResponse struct {
		UserID int64
	}
)

type (
	LoginUserRequest struct {
		Email    string
		Password string
	}

	LoginUserResponse struct {
		Token string
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
	TransactionInfoRequest struct {
		ChainName string
		TxHash    string
	}

	TransactionInfoResponse struct {
		BlockSignedAt     string
		BlockHeight       int64
		BlockHash         string
		TxHash            string
		TxOffset          int64
		Successful        bool
		FromAddress       string
		MinerAddress      string
		FromAddressLabel  string
		ToAddress         string
		ToAddressLabel    string
		Value             string
		ValueQuote        float64
		PrettyValueQuote  string
		GasMetadata       GasMetadataInfo
		GasOffered        int64
		GasSpent          int64
		GasPrice          int64
		FeesPaid          string
		GasQuote          float64
		PrettyGasQuote    string
		GasQuoteRate      float64
		Explorers         []ExplorerInfo
		LogEvents         []LogEventInfo
		InternalTransfers []InternalTransferInfo
		StateChanges      []StateChangeInfo
		InputData         InputDataInfo
	}

	GasMetadataInfo struct {
		ContractDecimals     int64
		ContractName         string
		ContractTickerSymbol string
		ContractAddress      string
		SupportsERC          []string
		LogoURL              string
	}

	ExplorerInfo struct {
		Label string
		URL   string
	}

	LogEventInfo struct {
		BlockSignedAt          string
		BlockHeight            int64
		TxOffset               int64
		LogOffset              int64
		TxHash                 string
		RawLogTopics           []string
		SenderContractDecimals int64
		SenderName             string
		SenderTickerSymbol     string
		SenderAddress          string
		SenderAddressLabel     string
		SenderLogoURL          string
		SupportsERC            []string
		SenderFactoryAddress   string
		RawLogData             string
		Decoded                DecodedDataInfo
	}

	DecodedDataInfo struct {
		Name      string
		Signature string
		Params    []ParamInfo
	}

	ParamInfo struct {
		Name    string
		Type    string
		Indexed bool
		Decoded bool
		Value   interface{}
	}

	InternalTransferInfo struct {
		FromAddress string
		ToAddress   string
		Value       string
		GasLimit    int64
	}

	StorageChangeInfo struct {
		StorageAddress string
		ValueBefore    string
		ValueAfter     string
	}

	StateChangeInfo struct {
		Address        string
		BalanceBefore  string
		BalanceAfter   string
		StorageChanges []StorageChangeInfo
		NonceBefore    int64
		NonceAfter     int64
	}

	InputDataInfo struct {
		MethodID string
	}
)

type (
	TokenBalancesRequest struct {
		WalletAddress string
		ChainName     string
	}

	TokenBalancesResponse struct {
		Address       string
		ChainID       int
		ChainName     string
		QuoteCurrency string
		UpdatedAt     string
		Items         []TokenBalanceItem
	}

	LogoURLs struct {
		TokenLogoURL    string
		ProtocolLogoURL string
		ChainLogoURL    string
	}

	ProtocolMetadata struct {
		ProtocolName string
	}

	ExternalData struct {
		Name         string
		Description  string
		Image        string
		Image256     string
		Image512     string
		Image1024    string
		AnimationURL string
		ExternalURL  string
		Owner        string
	}

	NFTData struct {
		TokenID           string
		TokenBalance      string
		TokenURL          string
		SupportsERC       []string
		TokenPriceWei     string
		TokenQuoteRateEth string
		OriginalOwner     string
		ExternalData      ExternalData
		Owner             string
		OwnerAddress      string
		Burned            bool
	}

	TokenBalanceItem struct {
		ContractDecimals     int
		ContractName         string
		ContractTickerSymbol string
		ContractAddress      string
		ContractDisplayName  string
		SupportsERC          []string
		LogoURL              string
		LogoURLs             LogoURLs
		LastTransferredAt    string
		NativeToken          bool
		Type                 string
		IsSpam               bool
		Balance              string
		Balance24h           string
		QuoteRate            int
		QuoteRate24h         int
		Quote                int
		Quote24h             int
		PrettyQuote          string
		PrettyQuote24h       string
		ProtocolMetadata     ProtocolMetadata
		NFTData              []NFTData
	}
)

type (
	HistoricalPortfolioValueRequest struct {
		WalletAddress string
		ChainName     string
	}

	PricePoint struct {
		Balance     string
		Quote       int
		PrettyQuote string
	}

	Holding struct {
		QuoteRate int
		Timestamp string
		Close     PricePoint
		High      PricePoint
		Low       PricePoint
		Open      PricePoint
	}

	HistoricalPortfolioItem struct {
		ContractAddress      string
		ContractDecimals     int
		ContractName         string
		ContractTickerSymbol string
		LogoURL              string
		Holdings             []Holding
	}

	HistoricalPortfolioValueResponse struct {
		Address       string
		UpdatedAt     string
		QuoteCurrency string
		ChainID       int
		ChainName     string
		Items         []HistoricalPortfolioItem
	}
)
