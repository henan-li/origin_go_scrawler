package parser

import (
	"../../engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
	)


func ParseCity(contents []byte) engine.ParseResult {

	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, v := range matches {
		name := string(v[2])

		result.Items = append(result.Items, "User "+name)

		result.Requests = append(result.Requests, engine.Request{
			Url:        string(v[1]),
			ParserFunc: func (c []byte) engine.ParseResult {
				return ParseProfile(c,name)
			},
		})
	}

	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests,engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}
	return result
}
