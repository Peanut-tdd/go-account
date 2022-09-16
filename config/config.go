package config

type Config struct {
	Mysql `yaml:"mysql" json:"mysql"`
	Redis `yaml:"redis" json:"redis"`
	Wx    `yaml:"wx" json:"wx"`
}
