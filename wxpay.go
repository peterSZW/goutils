package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type TenpayConfig struct {
	APPID      string
	APPKEY     string
	SIGNTYPE   string
	PARTNERKEY string
	APPSERCERT string
}

type TenpayRequest struct {
	BankType   string
	Body       string
	Partner    string
	OutTradeNo string

	TotalFee float64
	FeeType  string

	NotifyUrl     string
	SpbillCreatIp string
	InputCharset  string
}

type TenpayResponse struct {
	BuyerEmail  string
	OutTradeNo  string
	TradeStatus string
	Subject     string
	TotalFee    float64
}

func sha1sign(str string) string {
	h := sha1.New()
	io.WriteString(h, str)

	return fmt.Sprintf("%x", h.Sum(nil))
}

func wxmd5Sign(str, key string) string {
	h := md5.New()
	//h := sha1.New()
	io.WriteString(h, str)
	io.WriteString(h, key)

	return fmt.Sprintf("%x", h.Sum(nil))
}

func wx_demo() {
	c := TenpayConfig{
		APPID:      "wxf8b4f85f3a794e77",                                                                                                               //appid
		APPKEY:     "2Wozy2aksie1puXUBpWD8oZxiD1DfQuEaiC7KcRATv1Ino3mdopKaPGQQ7TtkNySuAmCaDCrw4xhPY5qKTBl7Fzm0RgR3c0WaVYIXZARsxzHV2x7iwPPzOz94dnwPWSn", //paysign key
		SIGNTYPE:   "sha1",                                                                                                                             //method
		PARTNERKEY: "8934e7d15453e97507ef794cf7b0519d",                                                                                                 //通加密串
		APPSERCERT: "09cb46090e586c724d52f7ec9e60c9f8",
	}
	r := TenpayRequest{
		BankType:   "WX",
		Body:       "test",
		Partner:    "1900000109",
		OutTradeNo: "JcHAYX3beZbc7djt",

		TotalFee: 1.01,
		FeeType:  "1",

		NotifyUrl:     "htttp://www.baidu.com",
		SpbillCreatIp: "127.0.0.1",
		InputCharset:  "GBK",
	}

	fmt.Println(WxpayNewPage(c, r, "1394605992", "xWxRCcfcuAM4TaYy"))
}

//func wxverifySign(c TenpayConfig, u url.Values) (err error) {
//	p := kvpairs{}
//	sign := ""
//	for k := range u {
//		v := u.Get(k)
//		switch k {
//		case "sign":
//			sign = v
//			continue
//		case "sign_type":
//			continue
//		}
//		p = append(p, kvpair{k, v})
//	}
//	if sign == "" {
//		err = fmt.Errorf("sign not found")
//		return
//	}
//	p = p.RemoveEmpty()
//	p.Sort()
//	fmt.Println(u)
//	if md5Sign(p.Join(), c.Key) != sign {
//		err = fmt.Errorf("sign invalid")
//		return
//	}
//	return
//}

//func WxpayParseResponse(c TenpayConfig, p url.Values) (r Response, err error) {
//	if err = verifySign(c, p); err != nil {
//		return
//	}
//	r.BuyerEmail = p.Get("buyer_email")
//	r.TradeStatus = p.Get("trade_status")
//	r.OutTradeNo = p.Get("out_trade_no")
//	r.Subject = p.Get("subject")
//	fmt.Sscanf(p.Get("total_fee"), "%f", &r.TotalFee)

//	if r.TradeStatus != "TRADE_SUCCESS" && r.TradeStatus != "TRADE_FINISHED" {
//		err = fmt.Errorf("trade not success or finnished")
//		return
//	}
//	return
//}

type WxpayJson struct {
	AppId     string `json:"appId"`
	Package   string `json:"package"`
	TimeStamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	PaySign   string `json:"paySign"`
	SignType  string `json:"signType"`
}

func WxpayNewPage(c TenpayConfig, r TenpayRequest, timestampstr string, mynoncestr string) string {
	p := kvpairs{
		kvpair{`bank_type`, r.BankType},
		kvpair{`body`, r.Body},
		kvpair{`partner`, r.Partner},
		kvpair{`out_trade_no`, r.OutTradeNo},
		kvpair{`total_fee`, fmt.Sprintf("%.2f", r.TotalFee)},

		kvpair{`fee_type`, r.FeeType},
		kvpair{`notify_url`, r.NotifyUrl},
		kvpair{`spbill_create_ip`, r.SpbillCreatIp},
		kvpair{`input_charset`, r.InputCharset},
	}
	p = p.RemoveEmpty()
	p.Sort()

	sign := strings.ToUpper(wxmd5Sign(p.Join()+"&key="+c.PARTNERKEY, ""))

	p = append(p, kvpair{`sign`, sign})

	//fmt.Println(p.Join())

	p2 := kvpairs{
		kvpair{`bank_type`, r.BankType},
		kvpair{`body`, r.Body},
		kvpair{`partner`, r.Partner},
		kvpair{`out_trade_no`, r.OutTradeNo},
		kvpair{`total_fee`, fmt.Sprintf("%.2f", r.TotalFee)},

		kvpair{`fee_type`, r.FeeType},
		kvpair{`notify_url`, r.NotifyUrl},
		kvpair{`spbill_create_ip`, r.SpbillCreatIp},
		kvpair{`input_charset`, r.InputCharset},
	}
	p2 = p2.RemoveEmpty()
	p2 = p2.UrlEncode()
	p2.Sort()
	p2 = append(p2, kvpair{`sign`, sign})

	ss := kvpairs{
		kvpair{`appid`, c.APPID},
		kvpair{`package`, p2.Join()},
		kvpair{`timestamp`, timestampstr},
		kvpair{`noncestr`, mynoncestr},
		kvpair{`appkey`, c.APPKEY},
	}
	ss = ss.RemoveEmpty()
	ss.Sort()

	//fmt.Println(sha1sign(ss.Join()))

	wxpay := WxpayJson{
		AppId:     c.APPID,
		Package:   "[!~~~~~~~~!]",
		TimeStamp: timestampstr,
		NonceStr:  mynoncestr,
		PaySign:   sha1sign(ss.Join()),
		SignType:  "sha1",
	}
	jsonstr, jsonerr := json.Marshal(wxpay)

	result := string(jsonstr)
	if jsonerr != nil {
		fmt.Println("error:", jsonerr)
		return ""
	} else {
		result = string(jsonstr)
	}
	result = strings.Replace(result, "[!~~~~~~~~!]", p2.Join(), 1)

	return result

}

func WxpayTesttenpay(Websiteroot string, accname string, ordernumber string, subject string, price float64, id string, key string) string {
	//r := TenpayRequest{
	//	NotifyUrl:   Websiteroot + `orderconfirm`,           // 付款后异步通知页面
	//	ReturnUrl:   Websiteroot + `orderreturn/` + accname, // 付款后返回页面
	//	OutTradeNo:  ordernumber,                            // 订单号
	//	SellerEmail: `sales@chat4support.com.cn`,            // 支付宝用户名
	//	Service:     `create_direct_pay_by_user`,            // 不可改
	//	PaymentType: `1`,                                    // 不可改
	//	Subject:     subject,                                // 商品名称
	//	TotalFee:    price,                                  // 价格
	//}

	// 输出的是 html 页面，会自动跳转到支付界面
	//return NewPage(c, r)
	return ""
}
