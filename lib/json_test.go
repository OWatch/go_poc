package lib

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonBuild(t *testing.T) {
	type TraceInfo struct {
		TraceId string `json:"trace_id"`
	}

	type Outer struct {
		TraceInfo *TraceInfo `json:"trace_info"`
	}
	type OuterWrapper struct {
		Outer *string `json:"outer"`
	}

	outer := &Outer{
		TraceInfo: &TraceInfo{
			TraceId: "test_id",
		},
	}

	outerByte, err := json.Marshal(outer)
	if err != nil {
		fmt.Printf("err:%+v", err)
	}

	outerStr := string(outerByte)

	outerWrapper := &OuterWrapper{
		Outer: &outerStr,
	}

	outerWrapperStr, _ := json.Marshal(outerWrapper)

	fmt.Printf("outerWrapperStr:%s", string(outerWrapperStr))
}
