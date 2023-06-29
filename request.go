package kmaparser

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)


type params struct {
	Tm string // 시간 (UTC), 없으면 가장 최근 데이터
	Stn string // 지점번호
	Mode string // L, H
	Help string // 0 or 1
	AuthKey string // 인증키
}

func NewParams(tm string, stn string, mode string, help bool) *params {
	helpNum := "0"
	if help {
		helpNum = "1"
	}
	return &params{
		Tm: tm,
		Stn: stn,
		Mode: mode,
		Help: helpNum,
	}
}

func (p *params) SetAuthKey(authKey string) {
	p.AuthKey = authKey
}

func GetWindProfilerRaderResp(p params) (string, error) {
	baseURL := "https://apihub.kma.go.kr/"
	path := "api/typ01/url/kma_wpf.php"
	params := url.Values{}
	params.Add("tm", p.Tm)
	params.Add("stn", string(p.Stn))
	params.Add("mode", p.Mode)
	params.Add("help", p.Help)
	params.Add("authKey", p.AuthKey)

	u, _ := url.ParseRequestURI(baseURL)
	u.Path = path
	u.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", u)

	resp, err := http.Get(urlStr)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}