package main

import (
	"fmt"
	"time"
)

type Rsp struct {
	code int32
}

const (
	cosntAbc = "abc"
)

func main() {
	//fmt.Println("hello world!")

	//str := ""
	//slice := make([]*string, 6)
	//slice = append(slice, &str)
	//
	//for i, v := range slice {
	//	println(i, v)
	//}
	//var code int32
	//rsp := &Rsp{}
	//strArr := make([]string, 3)
	//strArr[0] = "hello"
	//strArr[1] = "jack"
	//strArr[2] = "world"

	//result := make(map[string]string)
	//if result["goods_id"] == "" {
	//	fmt.Println(2)
	//	fmt.Println(result["goods_id"])
	//} else {
	//	fmt.Println(result)
	//}

	//actTime, _ := strconv.ParseInt("1663601621", 10, 64)
	//activeTime := time.Unix(actTime, 0).Format("2006-01-02 15:04:05.000")
	//str := "这是一个测试的标题UYUTEHUBG"

	//dateNow := time.Now().AddDate(0, 0, -1).Format("2006-01-02 15:04:05")
	//startTime := time.Date(dateNow.Year(), dateNow.Month(), dateNow.Day(), 0, 0, 0, 0,
	//	dateNow.Location()).Unix()

	//today := time.Now().Format("20060102")

	//startTime1 := time.Now().Add(-time.Hour * 1).Format("2006-01-02 15:04:05")
	//startTime2 := time.Now().AddDate(0, 0, -7).Format("2006-01-02 15:04:05")
	//
	//fmt.Println(startTime1)
	//fmt.Println(startTime2)

	//today := time.Now().Format("20060102")
	//todayInt, _ := strconv.ParseInt(today, 10, 32)
	//fmt.Println(todayInt)
	//
	//dateInt, _ := strconv.ParseInt("20221128", 10, 32)

	//fmt.Println(dateInt)
	////fmt.Println(time.Now().Weekday())
	//fmt.Println(todayInt - dateInt)
	//
	//// The leap year 2016 had 366 days.
	//t1 := time.Now()
	//t2 := Date(2022, 11, 1)
	//days := t1.Sub(t2).Hours() / 24
	//fmt.Println(days) // 366

	//date := time.Now()
	//format := "20060102"
	//today, _ := time.Parse(format, time.Now().Format(format))
	//fmt.Println(today)
	//
	//then, _ := time.Parse(format, "20221128")
	//fmt.Println(then)
	//diff := today.Sub(then)
	//diffInSec := int64(diff.Seconds())
	//diffInDay := int(diff.Hours() / 24)
	//
	////func Since(t Time) Duration
	////Since returns the time elapsed since t.
	////It is shorthand for time.Now().Sub(t).
	//
	////fmt.Println(diff.Hours())          // number of Hours
	////fmt.Println(diff.Nanoseconds())    // number of Nanoseconds
	////fmt.Println(int64(diff.Minutes())) // number of Minutes
	//fmt.Println(diffInSec) // number of Seconds
	//fmt.Println(diffInDay) // number of days

	//ocrWord := "a"
	//wordRegExp := regexp.MustCompile("\\[(.*?)]")
	//words := wordRegExp.FindAllStringSubmatch("这是ocr[尼玛]", -1)
	//for _, word := range words {
	//	ocrWord = word[1]
	//	fmt.Println(word[1])
	//}
	//
	//fmt.Println(ocrWord)
	//fmt.Println(fmt.Sprintf("%d", time.Now().Unix()))
	//ensembleStyleScoreMap := make(map[string]string)
	//err := json.Unmarshal([]byte("{\"v2\":\"0.10113770514726639\"}"), &ensembleStyleScoreMap)
	//if err == nil {
	//	fmt.Println(err)
	//}

	//aggregation2Nums := make(map[string]string)
	//aggregation2Nums["a"] += 12
	//aggregation2Nums["b"] += 12
	//aggregation2Nums["b"] += 1
	//aggregation2Nums["a"] = "123"
	var a = "abc"
	fmt.Println(a)
	fmt.Println("efg")
	//if aggregation2Nums["c"] == "" {
	//	fmt.Println("empty")
	//}

}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func test(strArr []string) string {
	result := ""
	if strArr == nil || len(strArr) == 0 {
		return ""
	}

	for i, tag := range strArr {
		if i == len(strArr)-1 {
			result += tag
			break
		}
		result += fmt.Sprintf("%s|", tag)
	}

	return result
}

// GetPreStr 获取前n个字
func GetPreStr(str string, i int) string {
	if len(str) == 0 {
		return str
	}
	runes := []rune(str)
	pre := str
	if len(runes) >= i {
		pre = string(runes[:i])
	}
	return fmt.Sprintf("《%s》", pre)
}
