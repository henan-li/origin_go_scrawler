package parser

import (
	"../../engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {

	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1) // 匹配所有,因此将得到:[0]完整的a标签, [1]url, [2]城市名字

	result := engine.ParseResult{}
	limit := 3
	for _, v := range matches {
		result.Items = append(result.Items, "City "+string(v[2])) // 此处就是城市名字,考虑到类型不确定,因此用空接口做元素的类型约束

		result.Requests = append(result.Requests, engine.Request{
			Url:        string(v[1]), // 每个城市的url
			ParserFunc: ParseCity,    // 访问url后对新页面要执行的提取信息的方法
		})

		limit--
		if limit == 0{
			break
		}
	}

	return result
}
