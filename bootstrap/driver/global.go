package driver

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GVA_VP *viper.Viper
	GVA_DB *gorm.DB
	R_DB   *redis.Client
)
