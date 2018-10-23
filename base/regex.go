package main

import (
	"regexp"
	
	"fmt"
)

const text = `My email is cc@gmain.com
this is ca@qq.com
`

func main(){
	re  := regexp.MustCompile(`[a-zA-Z0-9]+@.+\..+`)
	match := re.FindAllString(text,-1)
	fmt.Println(match)
}
