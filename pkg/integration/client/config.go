package client

type ConfigClient struct {
	BaseURL string `yaml:"get_activity_base_url" default:"api.covalenthq.com"`
}

//type CrossChainConfig struct {
//	BaseURL string `yaml:"get_activity_base_url" default:"api.covalenthq.com"`
//}
