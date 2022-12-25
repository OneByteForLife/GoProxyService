package models

import (
	"crypto/tls"
	"io"
	"net/http"
	"strconv"

	gojson "github.com/goccy/go-json"
	"github.com/sirupsen/logrus"
)

type ProxyHttpResponse struct {
	IP   string `json:"Ip"`
	Port int    `json:"Port"`
	Ping int    `json:"Ping"`
}

var proxyList []ProxyHttpResponse

func FindProxy(total string) ([]ProxyHttpResponse, string) {
	var data []ProxyHttpResponse

	if len(proxyList) == 0 {
		ParceProxy()
	}

	t, err := strconv.Atoi(total)
	if err != nil {
		logrus.Errorf("Err str to int - %s", err)
		return nil, "Err invalid query"
	}

	if t > len(proxyList) {
		return nil, "Query out of range"
	}

	for idx, val := range proxyList {
		data = append(data, val)
		proxyList = append(proxyList[:idx], proxyList[idx+1:]...)
		if idx == t-1 {
			break
		}
	}

	return data, "Success"
}

func ParceProxy() {
	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}

	resp, err := client.Get("https://www.proxyscan.io/api/proxy?type=http")
	if err != nil {
		logrus.Errorf("Err request proxy - %s", err)
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("Err read response body - %s", err)
		return
	}

	err = gojson.Unmarshal(data, &proxyList)
	if err != nil {
		logrus.Errorf("Err unmarshal body to struct - %s", err)
		return
	}
}
