package parser

import (
	"regexp"
	"../../engine"
)


const UserRe =  `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult{
	re := regexp.MustCompile(UserRe)
	match := re.FindAllSubmatch(contents,-1)
	result := engine.ParseResult{}
	for _, m := range match{
		userName := string(m[2])
		result.Items = append(result.Items, "User :" + userName)
		result.Requests = append(

			/*
			 如何把userName这个参数传到下一层？
			毫无疑问是要使用Reauest.ParserFunc这个，给这个方法加个参数？
			那该动就打了
			可以使用函数式编程的方式，再构造一个函数
			 */

			result.Requests, engine.Request{
				Url: 		string(m[1]),
				ParserFunc: func(c []byte) engine.ParseResult {
					return ParseProfile(c ,userName)
				},
			})
	}
	return result
}
