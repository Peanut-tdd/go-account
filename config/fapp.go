package config

type Fapp struct {
	Alipay `yaml:"alipay" json:"alipay"`
}

type Alipay struct {
	Appid        string `yaml:"appid" json:"appid"`
	AliPublicKey string `yaml:"aliPublicKey" json:"aliPublicKey"`
	PrivateKey   string `yaml:"privateKey" josn:"privateKey"`
}
