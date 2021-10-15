package index

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/encrypt"
	"github.com/gofrs/uuid"
	"github.com/iwind/TeaGo/Tea"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type ExpireConfig struct {
	Expire struct {
		On   bool   `yaml:"on" json:"on"`
		Code string `yaml:"code" json:"code"`
		Time string `yaml:"time" json:"time"` // 到期字符串
	} `yaml:"expire" json:"expire"`
}

func fileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}
func initWriteConfigFile() (*ExpireConfig, error) {
	conf := &ExpireConfig{}
	conf.Expire.On = true

	data, err := yaml.Marshal(conf)
	if err != nil {
		return conf, err
	}
	err = ioutil.WriteFile(Tea.ConfigFile("expire.yaml"), data, 0666)
	if err != nil {
		return conf, err
	}
	return conf, nil
}

// 读取当前服务到期配置
func LoadServerExpireConfig() (*ExpireConfig, error) {

	expireConfig := &ExpireConfig{}
	var err error
	configFile := Tea.ConfigFile("expire.yaml")
	if !fileExist(configFile) { //不存在 则copy,init
		expireConfig, err = initWriteConfigFile()
		if err != nil {
			return nil, err
		}
	} else {
		data, err := ioutil.ReadFile(configFile)
		if err != nil {
			return nil, err
		}
		err = yaml.Unmarshal(data, expireConfig)
		if err != nil {
			return nil, err
		}
	}

	if expireConfig.Expire.Time != "" {
		expireConfig.Expire.Time = string(encrypt.MagicKeyDecode([]byte(expireConfig.Expire.Time)))
	}
	if expireConfig.Expire.Code != "" {
		expireConfig.Expire.Code = string(encrypt.MagicKeyDecode([]byte(expireConfig.Expire.Code)))
	}
	return expireConfig, nil
}

// 保存当前服务配置
func writeServerConfig(expireConfig *ExpireConfig) error {
	//加密time
	if expireConfig.Expire.Time != "" {
		expireConfig.Expire.Time = string(encrypt.MagicKeyEncode([]byte(expireConfig.Expire.Time)))
	}
	//加密code
	if expireConfig.Expire.Code != "" {
		expireConfig.Expire.Code = string(encrypt.MagicKeyEncode([]byte(expireConfig.Expire.Code)))
	}

	data, err := yaml.Marshal(expireConfig)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(Tea.ConfigFile("expire.yaml"), data, 0666)
	if err != nil {
		return err
	}

	return nil
}

func CheckExpire() (string, bool, error) {

	expireConfig, err := LoadServerExpireConfig()
	if err != nil {
		return "", false, err
	}
	var expire bool
	var write bool
	var code string
	if !expireConfig.Expire.On {
		expire = false
	} else { //解析time并判断当前时间

		if expireConfig.Expire.Code == "" {
			write = true
			expire = true
			v4, _ := uuid.NewV4()
			expireConfig.Expire.Code = strings.ReplaceAll(v4.String(), "-", "")
		} else {
			if expireConfig.Expire.Time == "" {
				expire = true
			} else {
				//到期时间
				expireT, err := strconv.ParseInt(expireConfig.Expire.Time, 10, 64)
				if err != nil {
					return "", false, err
				}
				//当前时间 大于 到期时间 则表示到期
				expire = time.Now().Unix() > expireT
			}
		}
		code = expireConfig.Expire.Code
	}
	if write {
		err = writeServerConfig(expireConfig)
		if err != nil {
			return code, false, err
		}
	}
	return code, expire, err
}
