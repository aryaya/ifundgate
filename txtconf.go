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
	quote_url      string
	// withdrawalFee  float64
	// transferFee    float64
}

func readIniFile(fname string) (*txtConf, error) {
	iniConf, err := config.NewConfig("ini", fname)
	if err != nil {
		return nil, err
	}

	conf := &txtConf{}

	m, err := iniConf.GetSection("currencies")
	if err != nil {
		log.Println(err)
	}
	for k, _ := range m {
		conf.currencies = append(conf.currencies, k)
		break
	}

	m, err = iniConf.GetSection("domain")
	if err != nil {
		log.Println(err)
	}
	for k, _ := range m {
		conf.domain = k
		break
	}

	m, err = iniConf.GetSection("accounts")
	if err != nil {
		return nil, err
	}
	for k, _ := range m {
		conf.accounts = k
		break
	}

	m, err = iniConf.GetSection("hotwallets")
	if err != nil {
		log.Println(err)
	}
	for k, _ := range m {
		conf.hotwallets = append(conf.hotwallets, k)
	}

	m, err = iniConf.GetSection("federation_url")
	if err != nil {
		log.Println(err)
	}
	for k, _ := range m {
		conf.federation_url = k
		break
	}

	m, err = iniConf.GetSection("quote_url")
	if err != nil {
		log.Println(err)
	}
	for k, _ := range m {
		conf.quote_url = k
		break
	}

	return conf, nil
}
