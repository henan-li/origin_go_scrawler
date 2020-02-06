package parser

import (
	"../../engine"
	"../../model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄: </span>([\d]+)岁</td>`)
var eduRe = regexp.MustCompile(`<td><span class="label">学历: </span>([^<]+)</td>`)
var geoRe = regexp.MustCompile(`<td><span class="label">籍贯: </span>([^<]+)</td>`)
var idRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func ParseProfile(contents []byte, name string, url string) engine.ParseResult {

	profile := model.Profile{}

	profile.Name = name

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		profile.Age = age
	}

	education := extractString(contents, eduRe)
	if err != nil {
		profile.Education = education
	}

	geo := extractString(contents, geoRe)
	if err != nil {
		profile.Geo = geo
	}

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Id:      extractString([]byte(url),idRe),
				Type:    "zhenai",
				Payload: profile,
			},
		},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
