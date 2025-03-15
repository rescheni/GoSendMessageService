package yiyanAPI

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Sentence struct {
	Hitokoto string `json:"hitokoto"`
	From     string `json:"from"`
	FromWho  string `json:"from_who"`
	Length   int    `json:"length"`
}

// 1		 a  动画
// 2		 b 	漫画
// 4		 c	游戏
// 8		 d	文学
// 16	 	 e	原创
// 32	 	 f	来自网络
// 64	 	 g	其他
// 128	 	 h	影视
// 256	 	 i	诗词
// 512	 	 j	网易云
// 1024	 	 k	哲学
// 2048	 	 l	抖机灵
func GetSentence(sentenceType int) string {
	settype := ""
	if sentenceType&1 == 1 {
		settype += "&c=a"
	}
	if sentenceType&2 == 2 {
		settype += "&c=b"
	}
	if sentenceType&4 == 4 {
		settype += "&c=c"
	}
	if sentenceType&8 == 8 {
		settype += "&c=d"
	}
	if sentenceType&16 == 16 {
		settype += "&c=e"
	}
	if sentenceType&32 == 32 {
		settype += "&c=f"
	}
	if sentenceType&64 == 64 {
		settype += "&c=g"
	}
	if sentenceType&128 == 128 {
		settype += "&c=h"
	}
	if sentenceType&256 == 256 {
		settype += "&c=i"
	}
	if sentenceType&512 == 512 {
		settype += "&c=j"
	}
	if sentenceType&1024 == 1024 {
		settype += "&c=k"
	}
	if sentenceType&2048 == 2048 {
		settype += "&c=l"
	}

	req, err := http.NewRequest("GET", "https://v1.hitokoto.cn/?encode=json"+settype, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var sentence Sentence
	err = json.Unmarshal(body, &sentence)
	if err != nil {
		log.Fatal(err)
	}

	return sentence.Hitokoto + "------" + sentence.FromWho + "------" + sentence.From + "\n"
}
