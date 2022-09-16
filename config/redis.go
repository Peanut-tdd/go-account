package config

type Redis struct {
	Address  string `yaml:"address"   json:"address"`
	Password string `yaml:"password"  json:"password"`
	Database int    `yaml:"database"  json:"database"`
}
