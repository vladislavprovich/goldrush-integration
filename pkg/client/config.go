package client

type Config struct {
	CrossChain   CrossChainConfig
	Balances     BalancesConfig
	Transactions TransactionsConfig
	Security     SecurityConfig
	Bitcoin      BitcoinConfig
	Prices       PricesConfig
	Utility      UtilityConfig
	APIKey       string `json:"api_key" env:"API_KEY" envDefault:"cqt_rQ6yfw3Y8fCMKqJKmKf494ffWqK3"`
}

type CrossChainConfig struct {
	BaseURL      string `json:"cross_chain_base_url" env:"CROSS_CHAIN_BASE_URL" envDefault:"api.covalenthq.com"`
	Version      string `json:"cross_chain_version" env:"CROSS_CHAIN_VERSION" envDefault:"v1"`
	Address      string `json:"cross_chain_address" env:"CROSS_CHAIN_ADDRESS" envDefault:"address"`
	Activity     string `json:"cross_chain_activity" env:"CROSS_CHAIN_ACTIVITY" envDefault:"activity"`
	AllChains    string `json:"cross_chain_all_chains" env:"CROSS_CHAIN_ALL_CHAINS" envDefault:"allchains"`
	Balances     string `json:"cross_chain_balances" env:"CROSS_CHAIN_BALANCES" envDefault:"balances"`
	Transactions string `json:"cross_chain_transactions" env:"CROSS_CHAIN_TRANSACTIONS" envDefault:"transactions"`
}

type BalancesConfig struct {
	BaseURL            string `json:"balances_base_url" env:"BALANCES_BASE_URL" envDefault:"api.covalenthq.com"`
	Version            string `json:"balances_version" env:"BALANCES_VERSION" envDefault:"v1"`
	Address            string `json:"balances_address" env:"BALANCES_ADDRESS" envDefault:"address"`
	BalancesV2         string `json:"balances_v2" env:"BALANCES_V2" envDefault:"balances_v2"`
	HistoricalBalances string `json:"historical_balances" env:"HISTORICAL_BALANCES" envDefault:"historical_balances"`
	BalancesNative     string `json:"balances_native" env:"BALANCES_NATIVE" envDefault:"balances_native"`
	TransfersV2        string `json:"transfers_v2" env:"TRANSFERS_V2" envDefault:"transfers_v2"`
	PortfolioV2        string `json:"portfolio_v2" env:"PORTFOLIO_V2" envDefault:"portfolio_v2"`
	Token              string `json:"balances_token" env:"BALANCES_TOKEN" envDefault:"token"`
	TokenHoldersV2     string `json:"token_holders_v2" env:"TOKEN_HOLDERS_V2" envDefault:"token_holders_v2"`
}

type TransactionsConfig struct {
	BaseURL             string `json:"transactions_base_url" env:"TRANSACTIONS_BASE_URL" envDefault:"api.covalenthq.com"`
	Version             string `json:"transactions_version" env:"TRANSACTIONS_VERSION" envDefault:"v1"`
	Address             string `json:"transactions_address" env:"TRANSACTIONS_ADDRESS" envDefault:"address"`
	TransactionV2       string `json:"transaction_v2" env:"TRANSACTION_V2" envDefault:"transaction_v2"`
	TransactionsSummary string `json:"transactions_summary" env:"TRANSACTIONS_SUMMARY" envDefault:"transactions_summary"`
	Bulk                string `json:"transactions_bulk" env:"TRANSACTIONS_BULK" envDefault:"bulk"`
	Transactions        string `json:"transactions" env:"TRANSACTIONS" envDefault:"transactions"`
	TransactionsV3      string `json:"transactions_v3" env:"TRANSACTIONS_V3" envDefault:"transactions_v3"`
	Page                string `json:"transactions_page" env:"TRANSACTIONS_PAGE" envDefault:"page"`
	Block               string `json:"transactions_block" env:"TRANSACTIONS_BLOCK" envDefault:"block"`
	BlockHash           string `json:"transactions_block_hash" env:"TRANSACTIONS_BLOCK_HASH" envDefault:"block_hash"`
}

type SecurityConfig struct {
	BaseURL   string `json:"security_base_url" env:"SECURITY_BASE_URL" envDefault:"api.covalenthq.com"`
	Version   string `json:"security_version" env:"SECURITY_VERSION" envDefault:"v1"`
	Approvals string `json:"security_approvals" env:"SECURITY_APPROVALS" envDefault:"approvals"`
	Nft       string `json:"security_nft" env:"SECURITY_NFT" envDefault:"nft"`
}

type BitcoinConfig struct {
	BaseURL      string `json:"bitcoin_base_url" env:"BITCOIN_BASE_URL" envDefault:"api.covalenthq.com"`
	Version      string `json:"bitcoin_version" env:"BITCOIN_VERSION" envDefault:"v1"`
	Address      string `json:"bitcoin_address" env:"BITCOIN_ADDRESS" envDefault:"address"`
	BTCMainnet   string `json:"bitcoin_btc_mainnet" env:"BITCOIN_BTC_MAINNET" envDefault:"btc-mainnet"`
	HdWallets    string `json:"bitcoin_hd_wallets" env:"BITCOIN_HD_WALLETS" envDefault:"hd_wallets"`
	BalancesV2   string `json:"bitcoin_balances_v2" env:"BITCOIN_BALANCES_V2" envDefault:"balances_v2"`
	Cq           string `json:"bitcoin_cq" env:"BITCOIN_CQ" envDefault:"cq"`
	Covalent     string `json:"bitcoin_covalent" env:"BITCOIN_COVALENT" envDefault:"covalent"`
	App          string `json:"bitcoin_app" env:"BITCOIN_APP" envDefault:"app"`
	Bitcoin      string `json:"bitcoin_bitcoin" env:"BITCOIN_BITCOIN" envDefault:"bitcoin"`
	Transactions string `json:"bitcoin_transactions" env:"BITCOIN_TRANSACTIONS" envDefault:"transactions"`
}

type PricesConfig struct {
	BaseURL                 string `json:"prices_base_url" env:"PRICES_BASE_URL" envDefault:"api.covalenthq.com"`
	Version                 string `json:"prices_version" env:"PRICES_VERSION" envDefault:"v1"`
	Pricing                 string `json:"prices_pricing" env:"PRICES_PRICING" envDefault:"pricing"`
	HistoricalByAddressesV2 string `json:"prices_historical_by_addresses_v2" env:"PRICES_HISTORICAL_BY_ADDRESSES_V2" envDefault:"historical_by_addresses_v2"`
	SpotPrices              string `json:"prices_spot_prices" env:"PRICES_SPOT_PRICES" envDefault:"spot_prices"`
	Pools                   string `json:"prices_pools" env:"PRICES_POOLS" envDefault:"pools"`
	Event                   string `json:"prices_event" env:"PRICES_EVENT" envDefault:"event"`
	GasPrices               string `json:"prices_gas_prices" env:"PRICES_GAS_PRICES" envDefault:"gas_prices"`
}

type UtilityConfig struct {
	BaseURL        string `json:"utility_base_url" env:"UTILITY_BASE_URL" envDefault:"api.covalenthq.com"`
	Version        string `json:"utility_version" env:"UTILITY_VERSION" envDefault:"v1"`
	Events         string `json:"utility_events" env:"UTILITY_EVENTS" envDefault:"events"`
	Topics         string `json:"utility_topics" env:"UTILITY_TOPICS" envDefault:"topics"`
	Address        string `json:"utility_address" env:"UTILITY_ADDRESS" envDefault:"address"`
	BlockV2        string `json:"utility_block_v2" env:"UTILITY_BLOCK_V2" envDefault:"block_v2"`
	Chains         string `json:"utility_chains" env:"UTILITY_CHAINS" envDefault:"chains"`
	Status         string `json:"utility_status" env:"UTILITY_STATUS" envDefault:"status"`
	ResolveAddress string `json:"utility_resolve_address" env:"UTILITY_RESOLVE_ADDRESS" envDefault:"resolve_address"`
}
