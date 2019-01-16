package robot

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

func postHttp(url string, params interface{}) {
	
	jsonB, err := json.Marshal(params)
	if err != nil {
		logrus.WithError(err).Error("http post json marshal failed")
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonB))
	if err != nil {
		logrus.WithError(err).WithField(url, params).Error("PostHttp function create http(s) request failed")
		return
	}
	req.Header.Set("Content-Type", "application/json")
	
	client := &http.Client{}
	resp,err := client.Do(req)
	if err != nil {
		logrus.WithError(err).Error("do http post failed")
		return
	}
	b,err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		logrus.WithError(err).Error("do http post failed")
		return
	}
	logrus.Info(string(b))
}
type sinaJson struct {
	Title string `json:"title"`
	Content string `json:"content"`
	Cover string `json:"cover"`
	Summary string `json:"summary"`
	Text string `json:"text"`
	AccessToken string `json:"access_token"`
	
}
func PostSinaHeadline(title,content,cover,summary,text string){
	access_token := viper.GetString("SinaWeiboToken")
	p := sinaJson{title,content,cover,summary,text,access_token}
	postHttp("https://open.weibo.com/wiki/Toutiao/api",p)
}