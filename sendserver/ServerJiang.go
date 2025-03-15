package sendserver

import (
	basic "GoMessageService/Basic"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func ServerJiang(text string, desp string) {

	cfg := basic.LoadConfig()

	key := cfg.ServerJiang.ServerJiangKey

	ret := scSend(text, desp, key)
	fmt.Println(ret)
}

func scSend(text string, desp string, key string) string {
	data := url.Values{}
	data.Set("text", text)
	data.Set("desp", desp)

	// 根据 sendkey 是否以 "sctp" 开头决定 API 的 URL
	var apiUrl string
	if strings.HasPrefix(key, "sctp") {
		// 使用正则表达式提取数字部分
		re := regexp.MustCompile(`sctp(\d+)t`)
		matches := re.FindStringSubmatch(key)
		if len(matches) > 1 {
			num := matches[1]
			apiUrl = fmt.Sprintf("https://%s.push.ft07.com/send/%s.send", num, key)
		} else {
			return "Invalid sendkey format for sctp"
		}
	} else {
		apiUrl = fmt.Sprintf("https://sctapi.ftqq.com/%s.send", key)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return err.Error()
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}

	return string(body)
}
