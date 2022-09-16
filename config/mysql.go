package config

type Mysql struct {
	Host     string `yaml:"host"   json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json "password"`
	Database string `yaml:"database" json:"database"`
	Charset  string `yaml:"charset" json:"charset"`
	Timeout  string `yaml:"timeout json："timeout"`
}
