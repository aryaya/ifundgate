//
//
//

package main

import (
	"encoding/json"
	"github.com/astaxie/beego/context"
)

type FederationResp struct {
	Result         string      `json:"result"`
	Error          string      `json:"error,omitempty"`
	ErrorMessage   string      `json:"error_message,omitempty"`
	FederationJson *Federation `json:"federation_json,omitempty"`
}

type Currency struct {
	Currency string `json:"currency"`
	Issuer   string `json:"issuer,omitempty"`
}

type ExtraField struct {
	Required bool   `json:"required"`
	Hint     string `json:"hint"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Lable    string `json:"lable"`
}

type Federation struct {
	Domain      string       `json:"domain"`
	Currencies  []Currency   `json:"currencies,omitempty"`
	ExtraFields []ExtraField `json:"extra_fields,omitempty"`
	Destination string       `json:"destination"`
	Type        string       `json:"type"`
	QuoteUrl    string       `json:"quote_url"`
}

func federationErrorResp(msg string) *FederationResp {
	return &FederationResp{
		Result:       "error",
		Error:        "-1",
		ErrorMessage: msg,
	}
}

func federationSucessResp(conf *txtConf, destination string) *FederationResp {
	currencies := []Currency{}
	for _, c := range conf.currencies {
		currencies = append(currencies, Currency{c, conf.accounts})
	}

	contactField := ExtraField{
		Required: true,
		Hint:     "请留邮箱或者 QQ, 如果提现过程出现问题，客服将通过这个信息联系你",
		Type:     "text",
		Name:     "contact_info",
		Lable:    "联系方式",
	}

	fields := []ExtraField{}
	if destination == "z" {
		fields = []ExtraField{
			ExtraField{
				Required: true,
				Hint:     "大额提现请走银行卡，如欲走银行卡请输入 y@ripplecn.com",
				Type:     "text",
				Name:     "alipay_account",
				Lable:    "支付宝提现，请输入支付宝账户",
			},
			ExtraField{
				Required: true,
				Hint:     "请输入支付宝账户对应的真实姓名",
				Type:     "text",
				Name:     "full_name",
				Lable:    "姓名",
			},
			contactField,
		}
	} else if destination == "y" {
		fields = []ExtraField{
			ExtraField{
				Required: true,
				Hint:     "请填入银行的名称",
				Type:     "text",
				Name:     "bank_name",
				Lable:    "银行卡提现，请填入银行名称",
			},
			ExtraField{
				Required: true,
				Hint:     "请填入银行卡卡号",
				Type:     "text",
				Name:     "card_number",
				Lable:    "银行卡号",
			},
			ExtraField{
				Required: true,
				Hint:     "待提现的银行账户的姓名",
				Type:     "text",
				Name:     "full_name",
				Lable:    "姓名",
			},
			ExtraField{
				Required: true,
				Hint:     "大于等于五万的提现请填写开户行名称",
				Type:     "text",
				Name:     "opening_branch",
				Lable:    "开户行",
			},
			contactField,
		}
	}

	return &FederationResp{
		Result: "success",
		FederationJson: &Federation{
			Domain:      conf.domain,
			Destination: destination,
			Type:        "federation_record",
			QuoteUrl:    conf.quote_url,
			Currencies:  currencies,
			ExtraFields: fields,
		},
	}
}

func federation(ctx *context.Context, conf *txtConf) {
	typ := ctx.Request.URL.Query().Get("type")
	if typ != "federation" {
		resp := federationErrorResp("the query type must be federation")
		sendResp(resp, ctx)
		return
	}

	destination := ctx.Request.URL.Query().Get("destination")
	if destination != "z" && destination != "y" {
		resp := federationErrorResp("the query destination must be z or y")
		sendResp(resp, ctx)
		return
	}

	domain := ctx.Request.URL.Query().Get("domain")
	if domain != conf.domain {
		resp := federationErrorResp("the query domain must be " + conf.domain)
		sendResp(resp, ctx)
		return
	}

	resp := federationSucessResp(conf, destination)
	sendResp(resp, ctx)
}

func sendResp(resp interface{}, ctx *context.Context) error {
	b, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	ctx.ResponseWriter.Write(b)
	return nil
}
