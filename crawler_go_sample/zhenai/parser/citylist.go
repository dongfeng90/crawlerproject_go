package parser

import (
	"regexp"
	"../../engine"
)


const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]*)</a>`

func ParseCityList(contents []byte) engine.ParseResult{
	re := regexp.MustCompile(cityListRe)
	match := re.FindAllSubmatch(contents,-1)
	result := engine.ParseResult{}
	limit := 20
	for _, m := range match{
		result.Items = append(result.Items, "City :" + string(m[2]))
		result.Requests = append(
			result.Requests, engine.Request{
				Url: 		string(m[1]),
				ParserFunc: ParseCity,
			})

		limit --
		if limit == 0{
			break
		}
	}
	return result
}
