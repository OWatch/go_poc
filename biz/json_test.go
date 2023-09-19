package biz

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func TestArray(t *testing.T) {
	lowValueContentModelResultStr := `{"default":"xxx"}`
	var lowValueContentModelResultJson interface{}
	err := json.Unmarshal([]byte(lowValueContentModelResultStr), &lowValueContentModelResultJson)
	if err != nil {
		fmt.Println("error1")
	}
	modelProbeStr, ok := lowValueContentModelResultJson.(map[string]interface{})
	if !ok {
		fmt.Println("error2")
	}
	modelProbe, _ := strconv.ParseFloat(modelProbeStr["default"].(string), 64)
	if modelProbe > 0.5 {
		fmt.Println("modelProbe")
	} else {
		fmt.Println(modelProbeStr["default"])
	}

	fmt.Println("end")
}

func TestMarshal(t *testing.T) {
	rawStr := `{'0': 0.9984046816825867, '1': 0.0012030618963763118, '2': 0.0003922685864381492}`
	rawStr = strings.ReplaceAll(rawStr, "'", "\"")
	adStruct := &struct {
		Prob float64 `json:"1"`
	}{}
	if err := jsoniter.UnmarshalFromString(rawStr, adStruct); err != nil {
		fmt.Printf("unmarshal error:%+v", err)
	}
	fmt.Printf("adStruct:%+v", adStruct)
}
