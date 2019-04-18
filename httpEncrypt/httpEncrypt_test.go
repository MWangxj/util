package httpEncrypt

import (
	"encoding/json"
	"fmt"
	"github.com/MWangxj/util/base64"
	"github.com/MWangxj/util/byte2str"
	"github.com/MWangxj/util/commonEncrypt"
	"net/http"
	"net/url"
	`os`
	`strings`
	"testing"
)

func TestGet(t *testing.T) {
	app := NewApp("apollo", "apoq2rEGljmefWfP", "apoq2rEGljmesalt")
	mm := map[string]string{
		"key":  "test",
		"name": "guishan",
	}
	var (
		res []byte
		err error
	)
	if res, err = Do(app, HttpGet, "192.168.49.97:8080/demo", nil, mm); err != nil {
		t.Fail()
		return
	}
	fmt.Println(byte2str.BytesToString(res))
}

func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		app := NewApp("simcode", "adsgsag2rEGljmefWfP", "dfasfhasfhuiahufd")
		mm := map[string]string{
			"key":  "test",
			"name": "guishan",
		}
		Do(app, HttpPost, "192.168.49.97:8080/demo", nil, mm)
	}
}

func BenchmarkPost(b *testing.B) {
	for i := 0; i < b.N; i++ {
		app := NewApp("apollo", "apoq2rEGljmefWfP", "apoq2rEGljmesalt")
		mm := map[string]string{
			"key":  "test",
			"name": "guishan",
		}
		Do(app, HttpPost, "127.0.0.1:8080/test", nil, mm)
	}
}

type SimInfo struct {
	NetInfo
	CardProp
}

// NetInfo 网络信息
type NetInfo struct {
	Vendor  *int `json:"vendor"`
	Company *int `json:"company"`
}

type CardProp struct {
	Iccid  string `json:"iccid"`
	Msisdn string `json:"msisdn"`
	Imsi   string `json:"imsi"`
	Vendor int    `json:"vendor"`
}
type SimProdInfo struct {
	SimInfo
	ProdInfos []ProdInfo `json:"prodinfos,prod_infos"`
}
type ProdInfo struct {
	ProdId          string `json:"prodid,prod_id"`
	ProdName        string `json:"prodname,prod_name"`
	ProdInfo        string `json:"prodInfo,prod_info"`
	ProdInsteffTime string `json:"prodinstefftime,prod_inst_eff-time"`
	ProdInstExptime string `json:"prodinstexptime,prod_inst_exp_time"`
	ProdGprs
	Sign string `json:"sign"`
}
type ProdGprs struct {
	Total int `json:"total,gprsTotal"`
	Used  int `json:"used,gprsUsed"`
	Left  int `json:"left,gprsLeft"`
}

func TestPost(t *testing.T) {

	app := NewApp("simcode", "adsgsag2rEGljmefWfP", "dfasfhasfhuiahufd")
	//mm := map[string]string{
	//	"key":  "test",
	//	"name": "guishan",
	//}

	var mm struct {
		Param map[string]string `json:"param"`
		Index int               `json:"index"`
		EbUrl string            `json:"eb_url"`
		Info  []SimInfo         `json:"info"`
	}
	mm.Param = make(map[string]string)
	mm.Param["msisdn"] = "1064711527817"
	//mm.Param["queryDate"]="20181119"
	//mm.Param["card_info"] = "898607B2111790002183"
	//mm.Param["type"]="2"
	mm.EbUrl = "userstatusrealsingle"
	mm.Index = 2
	//sim := &SimInfo{}
	//sim.Msisdn = "1064724339193"
	//sim.Iccid = "898607B2111790002183"
	//sim.Imsi = "460041243302183"
	//mm.Info = []SimInfo{*sim}
	var (
		res []byte
		err error
	)
	head := map[string]string{
		"Api-Key": "simcode",
	}
	// 192.168.48.189:8080/v2/device/syncInfo"
	// 59.110.53.169
	if res, err = Do(app, HttpPost, "10.25.132.238:9001/v1/sim/userstatus", head, mm); err != nil {
		t.Fail()
		return
	}
	fmt.Println(byte2str.BytesToString(res))
}

func TestSimInfo(t *testing.T) {

	app := NewApp("simcode", "adsgsag2rEGljmefWfP", "dfasfhasfhuiahufd")
	//mm := map[string]string{
	//	"key":  "test",
	//	"name": "guishan",
	//}

	var mm struct {
		Param map[string]string `json:"param"`
		//Index int               `json:"index"`
		//EbUrl string            `json:"eb_url"`
	}
	mm.Param = make(map[string]string)
	//mm.Param["iccid"] = "89860404101800021501"
	//mm.Param["queryDate"]="20181119"
	mm.Param["msisdn"] = "1440406639982"
	var (
		res []byte
		err error
	)
	head := map[string]string{
		"Api-Key": "simcode",
	}
	// 192.168.48.189:8080/v2/device/syncInfo"
	// 59.110.53.169
	if res, err = Do(app, HttpPost, "simcode.dian.so/v1/sim/cardprod", head, mm); err != nil {
		t.Fail()
		return
	}
	fmt.Println(byte2str.BytesToString(res))
}

func readFile(fp string) (string, error) {
	f, err := os.Open(fp)
	defer f.Close()
	if err != nil {
		return "", err
	}
	str := ""
	buf := make([]byte, 1024)
	for {
		n, _ := f.Read(buf)
		if 0 == n {
			break
		}
		str += byte2str.BytesToString(buf[:n])
	}
	return str, nil
}

func TestSimImport(t *testing.T) {
	app := NewApp("simcode", "adsgsag2rEGljmefWfP", "dfasfhasfhuiahufd")
	//mm := map[string]string{
	//	"key":  "test",
	//	"name": "guishan",
	//}

	var cards = struct {
		Info []CardProp `json:"info"`
	}{}

	str, er := readFile("/Users/apple/Documents/sim201808002w.txt")
	if er != nil {
		t.Fail()
		return
	}
	strArr := strings.Split(str, ",")

	for _, s := range strArr {
		if len(s) != 13 {
			continue
		}
		card := CardProp{
			Msisdn: s,
			Vendor: 0,
		}
		cards.Info = append(cards.Info, card)
	}
	var (
		res []byte
		err error
	)
	head := map[string]string{
		"Api-Key": "simcode",
	}
	// 192.168.48.189:8080/v2/device/syncInfo"
	// 59.110.53.169
	if res, err = Do(app, HttpPost, "59.110.53.169:23333/v1/sim/import", head, cards); err != nil {
		t.Fail()
		return
	}
	fmt.Println(byte2str.BytesToString(res))
}

func TestSimGetList(t *testing.T) {
	app := NewApp("simcode", "adsgsag2rEGljmefWfP", "dfasfhasfhuiahufd")
	//mm := map[string]string{
	//	"key":  "test",
	//	"name": "guishan",
	//}

	var query = struct {
		Msisdn string
		//Start  int
		//End    int
	}{}
	//query.Start=0
	//query.End=20
	query.Msisdn="1064711527817"

	var (
		res []byte
		err error
	)
	head := map[string]string{
		"Api-Key": "simcode",
	}
	// 192.168.48.189:8080/v2/device/syncInfo"
	// 59.110.53.169
	if res, err = Do(app,HttpPost, "simcode.dian.so/v1/sim/userstatus", head, query); err != nil {
		t.Fail()
		return
	}

	//if res, err = Do(app,HttpGet, "10.45.147.5:9001/v1/sim/prods", head, query); err != nil {
	//	t.Fail()
	//	return
	//}
	fmt.Println(byte2str.BytesToString(res))
}

func listen() {
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {

		uri := "http://" + request.RemoteAddr + request.RequestURI

		fmt.Println(uri)
		urlR, err := url.Parse(uri)
		if err != nil {
			return
		}
		val, _ := url.ParseQuery(urlR.RawQuery)
		d := val.Get("data")
		s := val.Get("sign")
		fmt.Println(d, s)
		fmt.Println(fmt.Sprintf("%v", urlR))

		b := make([]byte, 1<<11)
		n, _ := request.Body.Read(b)
		b = b[:n]
		mm := make(map[string]string)
		json.Unmarshal(b, &mm)
		data, _ := base64.Base64UrlDecoding(mm["data"])
		t := commonEncrypt.VerifySign(request.URL.Path, mm["sign"], mm["ts"], mm["data"], mm["v"], request.Header.Get("token"), "apoq2rEGljmesalt")
		fmt.Println(t)
		m, err := commonEncrypt.Decrypt(data, "adsgsag2rEGljmefWfP")
		if err != nil {
			writer.Write(byte2str.StringToBytes(err.Error()))
			return
		}
		writer.Write(m)
	})
	http.ListenAndServe(":8080", nil)
}

func TestDO(t *testing.T) {

	data := "i5EjN81O7-vtA9KnrVWsxyZoKlmLUrSBgU-8yxxJaMuK_pXwf1aNUTf20m9B7FCdHYN61PwcR8j25Ir_VoFSy2XPbI29scT5Vma1o2fsIwYdWharXjB_cGvE7aV_O-DMS4ZRDYn0uEqwDsPARJeWE1Y-0UjR-mjuF_BCj1izoz3ANd4tONLsIsEi6jE1RElgWB1CG71a94EyuQH9ui2SJg"
	t.Log(len(data))
	b, err := base64.Base64UrlDecoding(data)
	if err != nil {
		t.Fail()
		return
	}
	b, err = commonEncrypt.Decrypt(b, "apoq2rEGljmefWfP")
	if err != nil {
		t.Fail()
		return
	}
	t.Log(byte2str.BytesToString(b))
}

func TestBase64(t *testing.T) {
	data := `{"password":"ba62addf26df0cd3","secKey":"007727178f52d397","cloudId":"b32c131869449025085688","deviceInfoId":"1261541","deviceNo":"869449025085688"}`
	str := base64.Base64UrlEncodeing(byte2str.StringToBytes(data))
	t.Log(str)

	u := "http://127.0.0.1:8080/test?data=" + str + "&sign=asdiufhaiu"
	urlR, err := url.Parse(u)
	if err != nil {
		t.Log(err)
		return
	}
	val, _ := url.ParseQuery(urlR.RawQuery)
	d := val.Get("data")
	s := val.Get("sign")
	fmt.Println(d, s)
	fmt.Println(fmt.Sprintf("%v", urlR))
}

func TestListen(t *testing.T) {
	listen()
}

func TestUrlParse(t *testing.T) {
	u := "http://127.0.0.1:8080/test?data=abc&sign=asdiufhaiu"
	urlR, err := url.Parse(u)
	if err != nil {
		t.Log(err)
		return
	}
	val, _ := url.ParseQuery(urlR.RawQuery)
	d := val.Get("data")
	s := val.Get("sign")
	fmt.Println(d, s)
	fmt.Println(fmt.Sprintf("%v", urlR))
}
