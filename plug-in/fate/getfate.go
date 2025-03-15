package fate

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/html"
)

// func main() {

// 	ret := FateTest()
// 	for _, v := range ret {
// 		fmt.Print(v)
// 	}
// }

func FateTest() []string {
	now := time.Now()
	year := now.Year()
	month := int(now.Month())
	day := now.Day()
	hour := now.Hour()

	if month == 1 && day == 1 {
		ret := make([]string, 0)
		ret = append(ret, GetYearRules(year)...)
		ret = append(ret, GetDailyRules(year, month, day, hour)...)
		ret = append(ret, GetMountRules(year, month)...)
		return ret
	} else if day == 1 {
		ret := make([]string, 0)
		ret = append(ret, GetMountRules(year, month)...)
		ret = append(ret, GetDailyRules(year, month, day, hour)...)
		return ret
	} else {
		ret := make([]string, 0)
		ret = append(ret, GetDailyRules(year, month, day, hour)...)
		return ret
	}
}

func GetDailyRules(year, month, day, hour int) []string {
	client := &http.Client{}
	var data = strings.NewReader(`FUNC=Basic&Name=&Solar=1&Year=2005&Month=8&Day=4&Hour=13&Sex=1&Target=5&SubTarget=0&Old=0&FateYearType=0&FateSolar=0&FateYear=` + strconv.Itoa(year) + `&FateMonth=` + strconv.Itoa(month) + `&FateDay=` + strconv.Itoa(day) + `&FateHour=` + strconv.Itoa(hour))
	req, err := http.NewRequest("POST", "https://fate.windada.com/cgi-bin/fate_gb", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en-US;q=0.8,en-GB;q=0.7,en;q=0.6")
	req.Header.Set("cache-control", "max-age=0")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	req.Header.Set("dnt", "1")
	req.Header.Set("origin", "https://fate.windada.com")
	req.Header.Set("priority", "u=0, i")
	req.Header.Set("referer", "https://fate.windada.com/cgi-bin/fate_gb")
	req.Header.Set("sec-ch-ua", `"Chromium";v="134", "Not:A-Brand";v="24", "Microsoft Edge";v="134"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36 Edg/134.0.0.0")
	req.Header.Set("cookie", "_ga=GA1.1.323049565.1740645109; V3bMale=1; V3bSolar=0; V3bYear=2005; V3bMonth=7; V3bDay=3; V3bHour=13; V2ShowLeap=0; FCNEC=%5B%5B%22AKsRol-CFyUTGJzbTSo5mMqK8Norjo93Dvt2xjGYB9CDf6tXWLYyznNp-ePCQ--mO2Te7pwvz4klYzVdOXMb2RZmPr6KK_Oe13R3qCijHIult8M7BSYz1UPJDxi7eoP_M8RdwkjcGgjotmMPT7mYMlrC0hij6bQpeg%3D%3D%22%5D%5D; _ga_JRD86XC779=GS1.1.1741856243.4.1.1741856883.59.0.0")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s\n", bodyText)

	ret := decode(string(bodyText))
	// for _, v := range ret {
	// 	fmt.Print(v)
	// }
	return ret

}

func decode(htmlContent string) []string {

	// 使用 html.Parse 函数解析 HTML 内容
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal(err)
	}

	//   定义变量保存需要提取的信息
	var 好运指数 string
	var 流日 string
	var 各星说明 []string

	// 定义递归函数来遍历 HTML 树
	var parseNode func(*html.Node)
	parseNode = func(n *html.Node) {

		// 查找包含 "好运指数" 的节点
		if n.Type == html.ElementNode && n.Data == "h2" {
			if n.FirstChild != nil && strings.Contains(n.FirstChild.Data, "好运指数") {
				好运指数 = n.FirstChild.Data
			}
		}

		// 查找包含 "流日" 的节点
		if n.Type == html.ElementNode && n.Data == "font" {
			if n.FirstChild != nil && strings.Contains(n.FirstChild.Data, "流日：") {
				流日 = n.FirstChild.Data
			}
		}

		// 查找 "各星说明" 的表格内容
		if n.Type == html.ElementNode && n.Data == "tbody" {
			// 遍历表格行，提取星说明
			for tr := n.FirstChild; tr != nil; tr = tr.NextSibling {
				if tr.Type == html.ElementNode && tr.Data == "tr" {
					// 这里查找每一行的星星和说明
					var starName, starDescription string
					for td := tr.FirstChild; td != nil; td = td.NextSibling {
						if td.Type == html.ElementNode && td.Data == "td" {
							if td.FirstChild != nil {
								if starName == "" {
									// 第一列是星星名称
									starName = td.FirstChild.Data
								} else {
									// 第二列是星星说明
									starDescription = td.FirstChild.Data
								}
							}
						}
					}
					// 如果都提取到了信息，保存
					if starName != "" && starDescription != "" {
						各星说明 = append(各星说明, fmt.Sprintf("%s：%s", starName, starDescription))
					}
				}
			}
		}

		// 递归遍历子节点
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			parseNode(c)
		}
	}

	// 开始解析文档
	parseNode(doc)

	ret := make([]string, 0)
	// 输出解析的结果
	ret = append(ret, fmt.Sprintln("好运指数:", 好运指数))
	ret = append(ret, fmt.Sprintln(流日))
	ret = append(ret, fmt.Sprintln("各星说明:"))
	flag := 0

	for _, star := range 各星说明 {
		if star == "table：table" {
			flag = 1
			continue
		}
		if flag == 1 {
			ret = append(ret, fmt.Sprintln(star))
		}
	}

	return ret

}

func GetMountRules(year, month int) []string {
	client := &http.Client{}
	var data = strings.NewReader(`FUNC=Basic&Name=&Solar=1&Year=2005&Month=8&Day=4&Hour=13&Sex=1&Target=4&SubTarget=0&Old=0&FateYearType=0&FateSolar=0&FateYear=` + strconv.Itoa(year) + `&FateMonth=` + strconv.Itoa(month) + `&FateDay=14&FateHour=0`)
	req, err := http.NewRequest("POST", "https://fate.windada.com/cgi-bin/fate_gb", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en-US;q=0.8,en-GB;q=0.7,en;q=0.6")
	req.Header.Set("cache-control", "max-age=0")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	req.Header.Set("dnt", "1")
	req.Header.Set("origin", "https://fate.windada.com")
	req.Header.Set("priority", "u=0, i")
	req.Header.Set("referer", "https://fate.windada.com/cgi-bin/fate_gb")
	req.Header.Set("sec-ch-ua", `"Chromium";v="134", "Not:A-Brand";v="24", "Microsoft Edge";v="134"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36 Edg/134.0.0.0")
	req.Header.Set("cookie", "_ga=GA1.1.323049565.1740645109; V3bMale=1; V3bSolar=0; V3bYear=2005; V3bMonth=7; V3bDay=3; V3bHour=13; V2ShowLeap=0; FCNEC=%5B%5B%22AKsRol-CFyUTGJzbTSo5mMqK8Norjo93Dvt2xjGYB9CDf6tXWLYyznNp-ePCQ--mO2Te7pwvz4klYzVdOXMb2RZmPr6KK_Oe13R3qCijHIult8M7BSYz1UPJDxi7eoP_M8RdwkjcGgjotmMPT7mYMlrC0hij6bQpeg%3D%3D%22%5D%5D; _ga_JRD86XC779=GS1.1.1741868455.5.1.1741868457.58.0.0")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	ret := decode(string(bodyText))
	// for _, v := range ret {
	// 	fmt.Print(v)
	// }
	return ret

}
func GetYearRules(year int) []string {
	client := &http.Client{}
	var data = strings.NewReader(`FUNC=Basic&Name=&Solar=1&Year=2005&Month=8&Day=4&Hour=13&Sex=1&Target=3&SubTarget=0&Old=0&FateYearType=0&FateSolar=0&FateYear=2025&FateMonth=3&FateDay=14&FateHour=0`)
	req, err := http.NewRequest("POST", "https://fate.windada.com/cgi-bin/fate_gb", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en-US;q=0.8,en-GB;q=0.7,en;q=0.6")
	req.Header.Set("cache-control", "max-age=0")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	req.Header.Set("dnt", "1")
	req.Header.Set("origin", "https://fate.windada.com")
	req.Header.Set("priority", "u=0, i")
	req.Header.Set("referer", "https://fate.windada.com/cgi-bin/fate_gb")
	req.Header.Set("sec-ch-ua", `"Chromium";v="134", "Not:A-Brand";v="24", "Microsoft Edge";v="134"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36 Edg/134.0.0.0")
	req.Header.Set("cookie", "_ga=GA1.1.323049565.1740645109; V3bMale=1; V3bSolar=0; V3bYear=2005; V3bMonth=7; V3bDay=3; V3bHour=13; V2ShowLeap=0; FCNEC=%5B%5B%22AKsRol-CFyUTGJzbTSo5mMqK8Norjo93Dvt2xjGYB9CDf6tXWLYyznNp-ePCQ--mO2Te7pwvz4klYzVdOXMb2RZmPr6KK_Oe13R3qCijHIult8M7BSYz1UPJDxi7eoP_M8RdwkjcGgjotmMPT7mYMlrC0hij6bQpeg%3D%3D%22%5D%5D; _ga_JRD86XC779=GS1.1.1741868455.5.1.1741869083.54.0.0")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	ret := decode(string(bodyText))
	// for _, v := range ret {
	// 	fmt.Print(v)
	// }
	return ret
}
