package config

type Config struct {
	Mysql  `yaml:"mysql" json:"mysql"`
	Redis  `yaml:"redis" json:"redis"`
	Wx     `yaml:"wx" json:"wx"`
	Server `yaml:"server" json:"server"`
}

type GetConfig interface {
	getKeyValue(key interface{}) interface{}
}

func (c *Config) GetKeyValue(item interface{}, key interface{}) (value interface{}) {



	switch item {

	case "Mysql":
		if key == nil {
			value = c.Mysql
			break
		}

		switch key {
		case "Host":
			value = c.Mysql.Host
			break
		}

	}

	return
}
