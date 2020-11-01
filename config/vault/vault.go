package vault

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	cfg "customer/config"

	"github.com/hashicorp/vault/api"
)

var (
	vaultClient *api.Client
)

type config struct {
	path string
	data map[string]interface{}
}

// GetConfig will get Vault address, token, and CA from environment variables.
// It will also get project environment from environment variable that will
// serve as the root path to get Vault secret.
//
// Project env environment variables should be in format of understandable
// Vault path, e.g cpt-dimii/dev/ hsn-remid/prd/
func GetConfig(path string) (cfg.Config, error) {
	address := os.Getenv("VAULT_ADDR")
	token := os.Getenv("VAULT_TOKEN")
	ca := os.Getenv("VAULT_CACERT")
	projectEnv := os.Getenv("PROJECT_ENV")
	return NewConfig(address, token, ca, projectEnv, path)
}

// NewConfig wii retrieve values from GetConfig function, then generate new config
// map that will serve as Vault client. Then, request to Vault will be sent with
// secret path generated from projectEnv and path variables as fullPath.
func NewConfig(address, token, ca, projectEnv, path string) (cfg.Config, error) {
	var err error
	vaultClient, err = createVaultClient(address, token, ca)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	fullPath := projectEnv + path
	resp, err := read(fullPath)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &config{path: fullPath, data: resp.Data}, nil
}

func createVaultClient(address, token, ca string) (*api.Client, error) {
	config := api.Config{
		Address: address,
	}
	if len(ca) > 0 {
		if err := config.ConfigureTLS(&api.TLSConfig{CACert: ca}); err != nil {
			log.Print(err)
			return nil, err
		}
	}

	client, err := api.NewClient(&config)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	client.SetToken(token)

	return client, nil
}

func read(path string) (*api.Secret, error) {
	resp, err := vaultClient.Logical().Read(path)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *config) GetInt(key string) int64 {
	value, ok := c.data[key]
	if ok {
		str := fmt.Sprintf("%s", value)
		num, err := strconv.ParseInt(str, 10, 64)
		if err == nil {
			log.Println("num err == nil: ", num)
			return num
		}
	}
	return 0
}

func (c *config) GetString(key string) string {
	value, ok := c.data[key]
	if ok {
		str, ok := value.(string)
		if ok {
			return str
		}
	}
	return ""
}

func (c *config) GetBool(key string) bool {
	value, ok := c.data[key]
	if ok {
		str, ok := value.(string)
		if ok {
			boolean, err := strconv.ParseBool(str)
			if err == nil {
				return boolean
			}
		}
	}
	return false
}
func (c *config) GetFloat(key string) float64 {
	value, ok := c.data[key]
	if ok {
		str, ok := value.(string)
		if ok {
			num, err := strconv.ParseFloat(str, 64)
			if err == nil {
				return num
			}
		}
	}
	return 0
}

func (c *config) GetBinary(key string) []byte {
	value, ok := c.data[key]
	if ok {
		str, ok := value.(string)
		if ok {
			bytes, err := base64.StdEncoding.DecodeString(str)
			if err == nil {
				return bytes
			}
		}
	}
	return nil
}

func (c *config) GetArray(key string) []string {
	value, ok := c.data[key]
	if ok {
		str, ok := value.(string)
		if ok {
			return strings.Split(str, ",")
		}
	}
	return nil
}

func (c *config) GetMap(key string) map[string]string {
	value, ok := c.data[key]
	if ok {
		str, ok := value.(string)
		if ok {
			maps := make(map[string]string)
			array := strings.Split(str, ",")
			for _, element := range array {
				kv := strings.SplitN(element, ":", 2)
				if len(kv) == 2 {
					maps[kv[0]] = kv[1]
				}
			}

			return maps
		}
	}
	return nil
}
