package getter

import (
	"encoding/json"
	"github.com/go-clog/clog"
	"github.com/henson/proxypool/pkg/models"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
)

type pubProxyReturn struct {
	Count int        `json:"count"`
	Data  []pubProxy `json:"data"`
}

type pubProxy struct {
	IpPort      string          `json:"ipPort"`
	Ip          string          `json:"ip"`
	Port        string          `json:"port"`
	Country     string          `json:"country"`
	LastChecked string          `json:"last_checked"`
	ProxyLevel  string          `json:"proxy_level"`
	Type        string          `json:"type"`
	Speed       string          `json:"speed"`
	Support     pubProxySupport `json:"support"`
}

type pubProxySupport struct {
	Https     int `json:"https"`
	Get       int `json:"get"`
	Post      int `json:"post"`
	Cookies   int `json:"cookies"`
	Referer   int `json:"referer"`
	UserAgent int `json:"user_agent"`
	Google    int `json:"google"`
}

func PubProxy() (result []*models.IP) {
	var ips pubProxyReturn
	var results []pubProxy

	pollURL := "http://pubproxy.com/api/proxy?https=true&limit=25&type=http&speed=25&last_check=60"
	resp, _, errs := gorequest.New().Get(pollURL).End()
	if errs != nil {
		clog.Error(0, "[PubProxy] error: %v", errs)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	err := json.Unmarshal(body, &ips)

	if err != nil {
		clog.Error(0, "[PubProxy] unmarshal body: %s, error: %s", string(body), err)
		return
	}

	results = ips.Data
	//fmt.Println("---->", ips.Count, ips.Data)
	for i := 0; i < len(results); i++ {
		ip := models.NewIP()
		ip.Data = results[i].IpPort
		ip.Type1 = "http"
		//ip.Type2 = "https"
		clog.Info("[PubProxy] ip.Data: %s,ip.Type: %s", ip.Data, ip.Type1)
		result = append(result, ip)
	}

	clog.Info("PubProxy done.")
	return
}
