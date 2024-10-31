package main

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

const qZoneSchemaPrefix = "xxx"

func testStringJoin() {
	stringArray1 := []string{"bob", "alice"}
	resultString1 := strings.Join(stringArray1, ",")
	fmt.Printf("resultString1:%s", resultString1)
	fmt.Println()

	stringArray2 := []string{"bob"}
	resultString2 := strings.Join(stringArray2, ",")
	fmt.Printf("resultString2:%s", resultString2)
	fmt.Println()

	var stringArray3 []string
	resultString3 := strings.Join(stringArray3, ",")
	fmt.Printf("resultString1:%s", resultString3)
	fmt.Println()
}

func testStringJoinAndMarshal() {
	extend := make(map[string]string, 0)
	titles := make([]string, 0)
	for i := 1; i < 11; i++ {
		title := "测试标题" + fmt.Sprintf("%d", i)
		titles = append(titles, title)
	}
	result, _ := jsoniter.MarshalToString(titles)
	extend["lucy_title"] = result
	extendStr, _ := jsoniter.MarshalToString(extend)
	fmt.Printf("extendStr:%s", extendStr)
	fmt.Println()

	lucyTitleArr := make([]string, 0)
	err := jsoniter.UnmarshalFromString(extend["lucy_title"], &lucyTitleArr)
	if err != nil {
		fmt.Printf("err extendStr:%s", result)
	}
	fmt.Printf("lucyTitleArr:%s", lucyTitleArr)
}

func testStringTrim() {
	str1 := "\"this is a str\""
	result1 := doTrim(str1)
	fmt.Printf("result1:%s", result1)
	fmt.Println()

	str2 := "this is a str"
	result2 := doTrim(str2)
	fmt.Printf("result2:%s", result2)
	fmt.Println()

	str3 := ""
	result3 := doTrim(str3)
	fmt.Printf("result3:%s", result3)
	fmt.Println()

	str4 := "\"\""
	result4 := doTrim(str4)
	fmt.Printf("result4:%s", result4)
	fmt.Println()

	trimResult := strings.TrimSuffix(strings.TrimPrefix("${material_type|10004}$", "${"), "}$")
	fmt.Printf("result5:%s", trimResult)
	fmt.Println()
	kvStrArr := strings.Split(trimResult, "|")
	if len(kvStrArr) > 1 {
		fmt.Printf("result6:%s", kvStrArr[1])
	}
}

func doTrim(rawStr string) string {
	if rawStr == "" {
		return rawStr
	}

	result := strings.TrimPrefix(rawStr, "\"")
	result = strings.TrimSuffix(result, "\"")

	return result
}

func AppendParameter(baseURL string, key string, value string) string {
	u, err := url.Parse(baseURL)
	if err != nil {
		return baseURL
	}

	queryParams := u.Query()
	queryParams.Set(key, value)

	u.RawQuery = queryParams.Encode()
	return u.String()
}

func foo(originJumpUrl string) {
	schemaEncoded := strings.TrimPrefix(originJumpUrl, qZoneSchemaPrefix)
	schemaDecodedBuffer, _ := base64.StdEncoding.DecodeString(schemaEncoded)
	originSchema := string(schemaDecodedBuffer)
	newSchema := AppendParameter(originSchema, "xxx", "xxx")
	newJumpUrl := fmt.Sprintf("%s%s", qZoneSchemaPrefix, base64.StdEncoding.EncodeToString([]byte(newSchema)))
	fmt.Printf("newJumpUrl2:%s", string(newJumpUrl))
	fmt.Println()
}
