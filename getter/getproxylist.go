package getter

import (
	"encoding/json"
	"fmt"
	"github.com/go-clog/clog"
	"github.com/henson/proxypool/pkg/models"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
)

type proxyReturn struct {
	Link                  map[string]interface{} `json:"link"`
	Ip                    string                 `json:"ip"`
	Port                  int                    `json:"port"`
	Protocol              string                 `json:"protocol"`
	Anonymity             string                 `json:"anonymity"`
	LastTested            string                 `json:"lastTested"`
	AllowsRefererHeader   bool                   `json:"allowsRefererHeader"`
	AllowsUserAgentHeader bool                   `json:"allowsUserAgentHeader"`
	AllowsCustomHeaders   bool                   `json:"allowsCustomHeaders"`
	AllowsCookies         bool                   `json:"allowsCookies"`
	AllowsPost            bool                   `json:"allowsPost"`
	AllowsHttps           bool                   `json:"allowsHttps"`
	Country               string                 `json:"country"`
	ConnectTime           string                 `json:"connectTime"`
	DownloadSpeed         string                 `json:"downloadSpeed"`
	SecondsToFirstByte    string                 `json:"secondsToFirstByte"`
	Uptime                string                 `json:"uptime"`
	Error                 string                 `json:"error,omitempty"`
}

func GetProxyList() (result []*models.IP) {
	var ips proxyReturn

	pollURL := "https://api.getproxylist.com/proxy?protocol[]=http&?lastTested=600&allowsHttps=1"
	resp, _, errs := gorequest.New().Get(pollURL).End()
	if errs != nil {
		clog.Error(0, "[GetProxyList] error: %v", errs)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	err := json.Unmarshal(body, &ips)

	if err != nil {
		clog.Error(0, "[GetProxyList] error: %s", err)
		return
	}
	if ips.Error != "" {
		clog.Error(0, "[GetProxyList] return error: %s", ips.Error)
		return
	}
	ip := models.NewIP()
	ip.Data = ips.Ip + ":" + fmt.Sprint(ips.Port)
	ip.Type1 = "http"
	ip.Type2 = "https"
	clog.Info("[GetProxyList] ip.Data: %s,ip.Type: %s,ip.Type2: %s", ip.Data, ip.Type1, ip.Type2)
	result = append(result, ip)
	clog.Info("GetProxyList done.")
	return
}
