package parser

import (
	"io/ioutil"
	"net/http"
	"pogo/common/logs"
)

type Xigua struct {
	Url string
}

func (xg *Xigua) GetVideoInfo() (*VideoInfo, error) {
	req, err := http.NewRequest(http.MethodGet, xg.Url, nil)
	if err != nil {
		return nil, err
	}
	header := http.Header{}
	header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4558.0 Safari/537.36 Edg/93.0.946.0")
	header.Add("referer", "https://www.ixigua.com/6984386884688577032")
	header.Add("upgrade-insecure-requests", "1")
	header.Add("sec-fetch-user", "?1")
	header.Add("sec-fetch-site", "same-origin")
	header.Add("sec-fetch-mode", "navigate")
	header.Add("sec-fetch-dest", "document")
	header.Add("sec-ch-ua-platform", "\"Windows\"")
	header.Add("sec-ch-ua-mobile", "?0")
	header.Add("sec-ch-ua", "\"Microsoft Edge\";v=\"93\", \" Not;A Brand\";v=\"99\", \"Chromium\";v=\"93\"")



	req.Header = header
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	resultBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	logs.Log.Debug(string(resultBytes))
	return &VideoInfo{}, nil
}
