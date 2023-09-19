package main

import (
	"bytes"
	"fmt"
	"text/template"
)

// NoticeRenderData 通知渲染数据
type NoticeRenderData struct {
	Content string
}

func main() {
	//renderData := &NoticeRenderData{
	//	Content: "bob",
	//}
	//result, err := renderContent(renderData, "this is {{.Content}}")
	//if err != nil {
	//	fmt.Printf("err:%+v", err)
	//}
	//fmt.Printf("result:%+v", result)

	//battleCount, err := strconv.ParseInt("", 10, 64)
	//if err != nil {
	//	fmt.Println("failed to parse ", err)
	//	return
	//}
	//fmt.Printf("result:%d", battleCount)
	//fmt.Println()

	type Person struct {
		Name string
	}

	bob := &Person{
		Name: "bob",
	}

	var persons []*Person
	persons = append(persons, bob)
	var arr []*Person = nil
	persons = append(persons, arr...)
	persons = append(persons, bob)
	fmt.Printf("len:%d", len(persons))
	fmt.Println()
	fmt.Printf("result:%+v", persons)
	fmt.Println()
}

func renderContent(renderData any, rawContent string) (string, error) {
	buf := &bytes.Buffer{}
	tmpl, err := template.New("tmpl").Parse(rawContent)
	if err != nil {
		fmt.Sprintf("failed to parse raw template %s, err %+v", rawContent, err)
		return "", err
	}
	if err := tmpl.Execute(buf, renderData); err != nil {
		fmt.Sprintf("failed to Execute raw template %s, err %+v", rawContent, err)
		return "", err
	}
	return buf.String(), nil
}
