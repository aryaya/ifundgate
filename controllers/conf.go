//
//
//

package controllers

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/wangch/glog"
)

// type GateBankAccount struct {
// 	Name       string   // 开户人姓名
// 	BankName   string   // 银行名字
// 	BankId     string   // 银行账号
// 	Currencies []string // 支持的货币种类
// 	Fees       Fees     // 交易费
// }

// type Fees struct {
// 	FeeMap map[string][2]float64 // 最低, 最高费率, 每笔转账小于最低按照最低计算, 高于最高按照最高计算
// 	Rate   float64               // 费率比率
// }

type Config struct {
	// GBAs       []GateBankAccount // 收款人信息
	Currencies []string // 支持的货币种类
	ColdWallet string   // 网关钱包地址 用于发行
	Host       string   // Server 地址
	// UsdRate    float64           // 当前 1 icc == ? usd 默认为1
	Domain string
}

var configFile = "./conf.json"

func loadConf() (*Config, error) {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	var conf Config
	err = json.Unmarshal(data, &conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}

// var defaultFees = Fees{
// 	FeeMap: map[string][2]float64{
// 		"CNY": {5, 50},
// 		"HKD": {6, 60},
// 		"USD": {1, 10},
// 		"EUR": {1, 10},
// 		"ICC": {1, 10},
// 		"BTC": {0.0005, 0.01},
// 	},
// 	Rate: 0.01,
// }

var defaultConf = &Config{
	Domain:     "isuncoin.com",
	Currencies: []string{"USD", "CNY", "HKD", "EUR", "JPY"},
	ColdWallet: "iN8sGowQCg1qptWcJG1WyTmymKX7y9cpmr", // ss1TCkz333t3t2J5eobcEMkMY3bXk // w01
	Host:       "localhost:8080",
}

func initConf() {
	conf, err := loadConf()
	if err != nil {
		conf = defaultConf
		b, err := json.MarshalIndent(conf, "", " ")
		if err != nil {
			glog.Fatal(err)
		}
		err = ioutil.WriteFile(configFile, b, os.ModePerm)
		if err != nil {
			glog.Fatal(err)
		}
	}
	Gconf = conf
}

var Gconf *Config

func init() {
	initConf()
}
