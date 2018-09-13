package parser

import (
	"testing"
	"../../fetcher"
	"fmt"
)

func TestParseCityList(t *testing.T) {
	contents ,err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	if err != nil{
		panic(err)
	}

	fmt.Printf("tst")
	result := ParseCityList(contents)

	const resultSize = 470

	if len(result.Requests) != resultSize{
		t.Errorf("error")
	}

	if len(result.Items) != resultSize{
		t.Errorf("error")
	}

	fmt.Println("success")

}