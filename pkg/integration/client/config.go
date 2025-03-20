package client

type Config struct {
	CrossChain   CrossChainConfig
	Balances     BalancesConfig
	Transactions TransactionsConfig
	Security     SecurityConfig
	Bitcoin      BitcoinConfig
	Prices       PricesConfig
	Utility      UtilityConfig
}

type CrossChainConfig struct {
	BaseURL      string `yaml:"cross_chain_base_url" default:"api.covalenthq.com"`
	Version      string `yaml:"cross_chain_version" default:"v1"`
	Address      string `yaml:"cross_chain_address" default:"address"`
	Activity     string `yaml:"cross_chain_activity" default:"activity"`
	AllChains    string `yaml:"cross_chain_all_chains" default:"allchains"`
	Balances     string `yaml:"cross_chain_balances" default:"balances"`
	Transactions string `yaml:"cross_chain_transactions" default:"transactions"`
}

type BalancesConfig struct {
	BaseURL            string `yaml:"cross_chain_base_url" default:"api.covalenthq.com"`
	Version            string `yaml:"cross_chain_version" default:"v1"`
	Address            string `yaml:"cross_chain_address" default:"address"`
	BalancesV2         string `yaml:"cross_chain_balances" default:"balances_v2"`
	HistoricalBalances string `yaml:"historical_balances" default:"historical_balances"`
	BalancesNative     string `yaml:"balances_native" default:"balances_native"`
	TransfersV2        string `yaml:"cross_chain_transfers" default:"transfers_v2"`
	PortfolioV2        string `yaml:"portfolio_v2" default:"portfolio_v2"`
	Token              string `yaml:"cross_chain_token" default:"token"`
	TokenHoldersV2     string `yaml:"token_holders_v2" default:"token_holders_v2"`
}

type TransactionsConfig struct {
	BaseURL             string `yaml:"cross_chain_base_url" default:"api.covalenthq.com"`
	Version             string `yaml:"cross_chain_version" default:"v1"`
	Address             string `yaml:"cross_chain_address" default:"address"`
	TransactionV2       string `yaml:"cross_chain_transactions" default:"transaction_v2"`
	TransactionsSummary string `yaml:"cross_chain_transactions_summary" default:"transactions_summary"`
	Bulk                string `yaml:"cross_chain_balances" default:"bulk"`
	Transactions        string `json:"cross_chain_transactions" default:"transactions"`
	TransactionsV3      string `json:"cross_chain_transactions_v3" default:"transactions_v3"`
	Page                string `yaml:"cross_chain_page" default:"page"`
	Block               string `yaml:"cross_chain_block" default:"block"`
	BlockHash           string `yaml:"cross_chain_block_hash" default:"block_hash"`
}

type SecurityConfig struct {
	BaseURL   string `yaml:"cross_chain_base_url" default:"api.covalenthq.com"`
	Version   string `yaml:"cross_chain_version" default:"v1"`
	Approvals string `yaml:"cross_chain_approvals" default:"approvals"`
	Nft       string `yaml:"cross_chain_nft" default:"nft"`
}

type BitcoinConfig struct {
	BaseURL      string `yaml:"cross_chain_base_url" default:"api.covalenthq.com"`
	Version      string `yaml:"cross_chain_version" default:"v1"`
	Address      string `yaml:"cross_chain_address" default:"address"`
	BTCMainnet   string `yaml:"cross_chain_btc_mainnet" default:"btc-mainnet"`
	HdWallets    string `yaml:"cross_chain_hd_wallets" default:"hd_wallets"`
	BalancesV2   string `yaml:"cross_chain_balances" default:"balances_v2"`
	Cq           string `yaml:"cross_chain_cq" default:"cq"`
	Covalent     string `yaml:"cross_chain_covalent" default:"covalent"`
	App          string `yaml:"cross_chain_app" default:"app"`
	Bitcoin      string `yaml:"cross_chain_bitcoin" default:"bitcoin"`
	Transactions string `yaml:"cross_chain_transactions" default:"transactions"`
}

type PricesConfig struct {
	BaseURL                 string `yaml:"cross_chain_base_url" default:"api.covalenthq.com"`
	Version                 string `yaml:"cross_chain_version" default:"v1"`
	Pricing                 string `yaml:"cross_chain_prices" default:"pricing"`
	HistoricalByAddressesV2 string `yaml:"historical_by_addresses_v_2" default:"historical_by_addresses_v2"`
	SpotPrices              string `yaml:"spot_prices" default:"spot_prices"`
	Pools                   string `yaml:"cross_chain_pools" default:"pools"`
	Event                   string `yaml:"cross_chain_event" default:"event"`
	GasPrices               string `yaml:"gas_prices" default:"gas_prices"`
}

type UtilityConfig struct {
	BaseURL        string `yaml:"cross_chain_base_url" default:"api.covalenthq.com"`
	Version        string `yaml:"cross_chain_version" default:"v1"`
	Events         string `yaml:"cross_chain_event" default:"events"`
	Topics         string `yaml:"cross_chain_topics" default:"topics"`
	Address        string `yaml:"cross_chain_address" default:"address"`
	BlockV2        string `yaml:"cross_chain_block" default:"block_v2"`
	Chains         string `yaml:"cross_chain_chains" default:"chains"`
	Status         string `yaml:"cross_chain_status" default:"status"`
	ResolveAddress string `yaml:"cross_chain_resolve_address" default:"resolve_address"`
