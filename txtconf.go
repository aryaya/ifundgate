//
//
//

package main

import (
	"github.com/astaxie/beego/config"
	"log"
)

type txtConf struct {
	currencies     []string
	domain         string
	accounts       string
	hotwallets     []string
	federation_url string
	withdrawalFee  float64
	transferFee    float64
}

func readIniFile(fname string) (*txtConf, error) {
	iniConf, err := config.NewConfig("ini", fname)
	if err != nil {
		return nil, err
	}

	conf := &txtConf{}

	m, err := iniConf.GetSection("currencies")
	if err != nil {
		return nil, err
	}
	for k, _ := range m {
		conf.currencies = append(conf.currencies, k)
	}

	m, err = iniConf.GetSection("domain")
	if err != nil {
		return nil, err
	}
	for k, _ := range m {
		conf.domain = k
	}

	m, err = iniConf.GetSection("accounts")
	if err != nil {
		return nil, err
	}
	for k, _ := range m {
		conf.accounts = k
	}

	m, err = iniConf.GetSection("hotwallets")
	if err != nil {
		log.Println(err)
	}
	for k, _ := range m {
		conf.hotwallets = append(conf.hotwallets, k)
	}

	conf.federation_url = iniConf.String("federation_url")

	conf.withdrawalFee, err = iniConf.Float("fees::withdrawal")
	if err != nil {
		log.Println("fees::withdrawal error:", err)
	}

	conf.transferFee, err = iniConf.Float("fees::transfer")
	if err != nil {
		log.Println("fees::transfer error:", err)
	}

	return conf, nil
}
