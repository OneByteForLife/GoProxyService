package models

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"
)

type ProxyHttpResponse struct {
	Type string `json:"type"`
	Data struct {
		IP   string `json:"ip"`
		Port string `json:"port"`
	} `json:"data"`
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
	c := colly.NewCollector()

	// За каждую итерацию 64
	for i := 64; i <= 640; {
		c.OnRequest(func(r *colly.Request) {
			r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
			r.Headers.Set("Cookie", "t=286995640; PAPVisitorId=3f29ea39b7fcb460b0ec429eca70c85Q; PAPVisitorId=3f29ea39b7fcb460b0ec429eca70c85Q; _ym_uid=1670355168894472826; _ym_d=1670355168; _tt_enable_cookie=1; _ttp=mHj0L660SSu4cKkgJySIbcunCif; cf_clearance=G9WY_em5uIgH1lMJUenX1DTWLuSO2SinCeU26kRCFYI-1671644155-0-150; _gid=GA1.2.1914327550.1671784608; _ym_isad=1; _ga_KJFZ3PJZP3=GS1.1.1671784608.5.1.1671786591.0.0.0; _ga=GA1.2.1961353494.1670355168; _gat_UA-90263203-1=1")
		})

		c.OnHTML("div > table tr", func(e *colly.HTMLElement) {
			ip := e.DOM.Find("td:nth-child(1)").Text()
			port := e.DOM.Find("td:nth-child(2)").Text()
			proxyType := e.DOM.Find("td:nth-child(5)").Text()
			if !strings.Contains(ip, "IP адрес") && !strings.Contains(port, "Порт") && !strings.Contains(proxyType, "Тип") {
				proxyList = append(proxyList, ProxyHttpResponse{
					Type: proxyType,
					Data: struct {
						IP   string "json:\"ip\""
						Port string "json:\"port\""
					}{
						IP:   ip,
						Port: port,
					},
				})
			}
		})
		// Требуется обновлять Cookie
		c.Visit(fmt.Sprintf("https://hidemy.name/ru/proxy-list/?maxtime=100&type=h&anon=1&start=%s#list", strconv.Itoa(i)))
		i += 64
	}
	fmt.Println(proxyList)
}
