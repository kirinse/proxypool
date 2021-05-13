package getter

import (
	"github.com/go-clog/clog"
	"github.com/henson/proxypool/pkg/models"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
	"strings"
)

func ProxyListDownload() (result []*models.IP) {
	pollURL := "https://www.proxy-list.download/api/v1/get?type=https&anon=elite"
	resp, _, errs := gorequest.New().Get(pollURL).End()
	if errs != nil {
		clog.Error(0, "[ProxyListDownload] error: %v", errs)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	lines := strings.Split(string(body), "\r\n")
	for _, line := range lines {
		data := strings.TrimSpace(line)
		if data == "" {
			continue
		}
		ip := models.NewIP()
		ip.Data = data
		ip.Type1 = "http"
		//ip.Type2 = "https"
		clog.Info("[ProxyListDownload] ip.Data: %s,ip.Type: %s", ip.Data, ip.Type1)
		result = append(result, ip)
	}
	return
}