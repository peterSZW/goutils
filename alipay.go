package goutils

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/url"

	"github.com/peterSZW/goutils/logger"
)

var alipayGatewayNew = `https://mapi.alipay.com/gateway.do?`

type Config struct {
	Partner string
	Key     string
}

type Request struct {
	Service     string
	PaymentType string
	NotifyUrl   string
	ReturnUrl   string
	OutTradeNo  string
	Subject     string
	TotalFee    float64
	Body        string
	ShowUrl     string
	SellerEmail string
}

type Response struct {
	BuyerEmail  string
	OutTradeNo  string
	TradeStatus string
	Subject     string
	TotalFee    float64
}

func md5Sign(str, key string) string {
	h := md5.New()
	io.WriteString(h, str)
	io.WriteString(h, key)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func verifySign(c Config, u url.Values) (err error) {
	p := kvpairs{}
	sign := ""
	for k := range u {
		v := u.Get(k)
		switch k {
		case "sign":
			sign = v
			continue
		case "sign_type":
			continue
		}
		p = append(p, kvpair{k, v})
	}
	if sign == "" {
		err = fmt.Errorf("sign not found")
		return
	}
	p = p.RemoveEmpty()
	p.Sort()
	logger.Debug(u)
	if md5Sign(p.Join(), c.Key) != sign {
		err = fmt.Errorf("sign invalid")
		return
	}
	return
}

func ParseResponse(c Config, p url.Values) (r Response, err error) {
	if err = verifySign(c, p); err != nil {
		return
	}
	r.BuyerEmail = p.Get("buyer_email")
	r.TradeStatus = p.Get("trade_status")
	r.OutTradeNo = p.Get("out_trade_no")
	r.Subject = p.Get("subject")
	fmt.Sscanf(p.Get("total_fee"), "%f", &r.TotalFee)

	if r.TradeStatus != "TRADE_SUCCESS" && r.TradeStatus != "TRADE_FINISHED" {
		err = fmt.Errorf("trade not success or finnished")
		return
	}
	return
}

func NewPage(c Config, r Request) string {
	p := kvpairs{
		kvpair{`_input_charset`, `utf-8`},
		kvpair{`out_trade_no`, r.OutTradeNo},
		kvpair{`partner`, c.Partner},
		kvpair{`payment_type`, r.PaymentType},
		kvpair{`notify_url`, r.NotifyUrl},
		kvpair{`return_url`, r.ReturnUrl},
		kvpair{`subject`, r.Subject},
		kvpair{`total_fee`, fmt.Sprintf("%.2f", r.TotalFee)},
		kvpair{`body`, r.Body},
		kvpair{`service`, r.Service},
		kvpair{`show_url`, r.ShowUrl},
		kvpair{`seller_email`, r.SellerEmail},
	}
	p = p.RemoveEmpty()
	p.Sort()

	sign := md5Sign(p.Join(), c.Key)
	p = append(p, kvpair{`sign`, sign})
	p = append(p, kvpair{`sign_type`, `MD5`})

	var result string
	result = ""
	result = result + `<html><head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	</head><body>`

	result = result + `<form name='alipaysubmit' action="` + alipayGatewayNew + `_input_charset=utf-8" method='post'> `
	for _, kv := range p {
		result = result + `<input type='hidden' name="` + kv.k + `" value="` + kv.v + `" />`
	}
	result = result + `<script>document.forms['alipaysubmit'].submit();</script>`
	result = result + `</body></html>`
	return result

}

//TestAlipay("12345","购物车",0.01)
func TestAlipay(Websiteroot string, accname string, ordernumber string, subject string, price float64, id string, key string) string {
	r := Request{
		NotifyUrl:   Websiteroot + `orderconfirm`,           // 付款后异步通知页面
		ReturnUrl:   Websiteroot + `orderreturn/` + accname, // 付款后返回页面
		OutTradeNo:  ordernumber,                            // 订单号
		SellerEmail: `sales@chat4support.com.cn`,            // 支付宝用户名
		Service:     `create_direct_pay_by_user`,            // 不可改
		PaymentType: `1`,                                    // 不可改
		Subject:     subject,                                // 商品名称
		TotalFee:    price,                                  // 价格
	}

	c := Config{
		Partner: id,  // 支付宝合作者身份 ID  `2088101107399755`
		Key:     key, // 支付宝交易安全校验码 `hvommt84ffgl9ljv57li1r4qg7nzq8o1`
	}

	// 输出的是 html 页面，会自动跳转到支付界面
	return NewPage(c, r)
}
