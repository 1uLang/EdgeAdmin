package index

import (
	"encoding/base64"
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/TeaOSLab/EdgeAdmin/internal/encrypt"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
	"strconv"
	"strings"
	"time"
)

// 检查是否需要OTP
type RenewalAction struct {
	actionutils.ParentAction
}

func (this *RenewalAction) Init() {
	this.Nav("", "", "")
}
func (this *RenewalAction) RunPost(params struct {
	Secret string

	Must *actions.Must
}) {
	params.Must.Field("secret", params.Secret).Require("请输入续订密钥")

	err := Renew(params.Secret)
	if err != nil {
		// 日志
		this.CreateLogInfo("系统续订失败:%v", err)
		this.ErrorPage(err)
	}
	this.Success()
}

// Renew secret ： code,expireTm,timeout,authNum
func Renew(secret string) error {

	//判断redis 是否存在该secret
	if v, _ := cache.GetCache(secret); v != nil && v != "" {
		return fmt.Errorf("密钥已失效，请不要重复续订")
	}
	decode, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return err
	}
	sec := encrypt.MagicKeyDecode(decode)
	//code,add_time,time
	secrets := strings.Split(string(sec), ",")
	if len(secrets) != 3 && len(secrets) != 4 {
		return fmt.Errorf("无效的密钥")
	}
	config, err := LoadServerExpireConfig()
	if err != nil {
		return err
	}
	//未开启 直接通过
	if !config.Expire.On {
		return nil
	}
	if config.Expire.Code != secrets[0] {
		return fmt.Errorf("不支持此密钥")
	}
	timeout, _ := strconv.ParseInt(secrets[2], 10, 64)
	if timeout < time.Now().Unix() {
		return fmt.Errorf("密钥已失效")
	}
	renewalT, _ := strconv.ParseInt(secrets[1], 10, 64)
	//新版授权域名个数
	if len(secrets) == 4 {
		num, _ := strconv.ParseInt(config.Expire.Num, 10, 64)
		addNum, _ := strconv.ParseInt(secrets[3], 10, 64)
		config.Expire.Num = fmt.Sprintf("%d", num+addNum)
	}
	//续订生效
	if config.Expire.Time == "" {
		expireT := time.Now().Unix() + renewalT
		config.Expire.Time = fmt.Sprintf("%v", expireT)
	} else {
		expireT, err := strconv.ParseInt(config.Expire.Time, 10, 64)
		if err != nil {
			return err
		}
		config.Expire.Time = fmt.Sprintf("%v", expireT+renewalT)
	}
	err = writeServerConfig(config)
	if err != nil {
		return err
	}
	//将secret 写入数据库防止重复续订
	_ = cache.SetCache(secret, true, uint32(timeout-time.Now().Unix()))
	return nil
}
