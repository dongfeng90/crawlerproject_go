package parser

import (
	"regexp"
	"../../engine"
	"../../model"
	"strconv"
)

var ageRe      = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var heighRe    = regexp.MustCompile(`<td><span class="label">身高：</span>[^<0-9]*([\d]+)CM[^<]*</td>`)
var weighRe    = regexp.MustCompile(`<td><span class="label">体重：</span>[^0-9]*([\d]+)KG`)
var sexRe      = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<])</span></td>`)
var incomeRe   = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var eduRe      = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occRe      = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var hokouRe    = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^>]+)</td>`)
var houseRe    = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe      = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)

func ParseProfile(contents []byte, name string) engine.ParseResult{
	profile := model.Profile{}
	profile.Age 	   = extractInt(contents, ageRe)
	profile.Height     = extractInt(contents, heighRe)
	profile.Weight     = extractInt(contents, weighRe)
	profile.Name       = name
	profile.Gender     = extractString(contents, sexRe)
	profile.Income     = extractString(contents, incomeRe)
	profile.Marriage   = extractString(contents, marriageRe)
	profile.Education  = extractString(contents, eduRe)
	profile.Occupation = extractString(contents, occRe)
	profile.HoKou      = extractString(contents, hokouRe)
	profile.House      = extractString(contents, houseRe)
	profile.Car        = extractString(contents, carRe)

	return engine.ParseResult{
		Items:[]interface{}{profile},
	}
}

func extractInt(contents [] byte, re *regexp.Regexp) int {
	age ,err := strconv.Atoi(extractString(contents, re))
	if err != nil{
		return -1
	}else{
		return age
	}
}


func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}else {
		return  ""
	}
}