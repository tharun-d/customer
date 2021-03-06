package viper

import (
	cfg "customer/config"
	"encoding/base64"
	"log"
	"strings"

	"github.com/spf13/viper"
)

type config struct {
}

func NewConfig(envPrefix, file string) (cfg.Config, error) {
	v := &config{}
	if err := v.init(envPrefix, file); err != nil {
		return nil, err
	}
	return v, nil
}

func (v *config) init(envPrefix, file string) error {
	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv()

	replacer := strings.NewReplacer(`.`, `_`)
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigType(`json`)
	viper.SetConfigFile(file)
	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (v *config) GetString(key string) string {
	return viper.GetString(key)
}

func (v *config) GetBool(key string) bool {
	return viper.GetBool(key)
}

func (v *config) GetInt(key string) int64 {
	return viper.GetInt64(key)
}

func (v *config) GetFloat(key string) float64 {
	return viper.GetFloat64(key)
}

func (v *config) GetBinary(key string) []byte {
	value := viper.GetString(key)
	bytes, err := base64.StdEncoding.DecodeString(value)
	if err == nil {
		return bytes
	}
	return nil
}

func (v *config) GetArray(key string) []string {
	return viper.GetStringSlice(key)
}

func (v *config) GetMap(key string) map[string]string {
	return viper.GetStringMapString(key)
}
