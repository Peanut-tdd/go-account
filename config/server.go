package config

type Server struct {
	Host string `yaml:"host",json:"host"`
	Port int    `yaml:"port" json:"port"`
}


