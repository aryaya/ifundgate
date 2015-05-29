//
//
//

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/astaxie/beego/context"
	"log"
	"strings"
	"time"
)

type QuoteResp struct {
	Result       string `json:"result"`
	Error        string `json:"error,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
	QuoteJson    *Quote `json:"quote,omitempty"`
}

type Amount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
	Issuer   string `json:"issuer,omitempty"`
}

func quoteErrorResp(msg string) *QuoteResp {
	return &QuoteResp{
		Result:       "error",
		Error:        "-1",
		ErrorMessage: msg,
	}
}

type Quote struct {
	Destination        string   `json:"destination,omitempty"`
	Type               string   `json:"type"`
	Domain             string   `json:"domain"`
	Amount             string   `json:"amount"`
	Source             string   `json:"Source"`
	DestinationAddress string   `json:"destination_address,omitempty"`
	DestinationTag     uint     `json:"destination_tag"`
	InvoiceID          string   `json:"invoice_id"`
	Send               []Amount `json:"send"`
	Expires            uint     `json:"expires"`
}

func getInvoiceID(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	hash := sha256.Sum256(b)
	invoiceID := hex.EncodeToString(hash[:])
	return invoiceID
}

func quoteSucessResp(ctx *context.Context, conf *txtConf, destination, domain, amount, source string) *QuoteResp {
	cAndV := strings.Split(amount, "/")
	if len(cAndV) != 2 {
		log.Fatal("amount is error")
	}

	send := []Amount{
		Amount{
			Currency: cAndV[1],
			Value:    cAndV[0],
			Issuer:   conf.accounts,
		},
	}

	quote := &Quote{
		Destination:        destination,
		Type:               "quote",
		Domain:             domain,
		Amount:             amount,
		Source:             source,
		DestinationAddress: conf.accounts,
		DestinationTag:     2147483647,
		Send:               send,
		InvoiceID:          getInvoiceID(send),
		Expires:            uint(time.Now().Unix() + 200),
	}
	return &QuoteResp{
		Result:    "success",
		QuoteJson: quote,
	}
}

func quote(ctx *context.Context, conf *txtConf) {
	typ := ctx.Request.URL.Query().Get("type")
	if typ != "quote" {
		resp := quoteErrorResp("the query type must be quote")
		sendResp(resp, ctx)
		return
	}

	domain := ctx.Request.URL.Query().Get("domain")
	if domain != conf.domain {
		resp := quoteErrorResp("the query domain must be " + conf.domain)
		sendResp(resp, ctx)
		return
	}

	destination := ctx.Request.URL.Query().Get("destination")
	if destination == "" {
		resp := quoteErrorResp("the query destination must be not null")
		sendResp(resp, ctx)
		return
	}

	amount := ctx.Request.URL.Query().Get("amount")
	if amount == "" || len(strings.Split(amount, "/")) != 2 {
		resp := quoteErrorResp("the query amount must be not null and must 1.0/CNY")
		sendResp(resp, ctx)
		return
	}

	source := ctx.Request.URL.Query().Get("source")
	if source == "" {
		resp := quoteErrorResp("the query amount must be not null")
		sendResp(resp, ctx)
		return
	}

	resp := quoteSucessResp(ctx, conf, destination, domain, amount, source)
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
