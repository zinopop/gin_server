package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)


type MsgContext struct {
	At struct {
		AtMobiles []string `json:"atMobiles"`
		IsAtAll   bool     `json:"isAtAll"`
	} `json:"at"`
	Msgtype string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}



func urlBuild() string {
	dingUrl := Config.GetString("DingTalk.SendUrl")+"?access_token="+Config.GetString("DingTalk.AccessToken")+"&timestamp="
	timestamp := time.Now().UnixNano() / 1e6
	timestampString :=strconv.FormatInt(timestamp,10)
	dingUrl += timestampString
	sign := timestampString + "\n" + Config.GetString("DingTalk.Secret")
	sign = ComputeHmacSha256(sign,Config.GetString("DingTalk.Secret"))
	sign = url.QueryEscape(sign)
	dingUrl += "&sign="+sign
	return dingUrl
}


func SendDingMsg(msg string) (string,error) {
	req, err := http.NewRequest("POST", urlBuild(), strings.NewReader(msg))
	if err != nil {
		return "请求错误1", err
	}
	client := &http.Client{}
	//设置请求头
	req.Header.Set("Content-Type", "application/json")
	//发送请求
	resp, err := client.Do(req)

	if err != nil{
		return "请求错误2", err
	}
	//关闭请求
	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "读取流失败", err
	}
	return string(body),nil
}

func ContextPush(msg string) []byte{
	re := &MsgContext{}
	re.Text.Content = msg
	re.Msgtype = "text"
	jsonBytes, err := json.Marshal(&re)
	if err != nil {
		fmt.Println(err)
	}
	return jsonBytes
}