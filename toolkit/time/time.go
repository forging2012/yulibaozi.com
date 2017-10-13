package time

import (
	"strconv"
	"time"
)

func ZhTime(t time.Time) string {
	var nowTimeSecond = time.Now().UnixNano() / int64(time.Second)
	var createTimeSecond = t.UnixNano() / int64(time.Second)
	diff := nowTimeSecond - createTimeSecond
	zhTime := "刚刚"
	if diff > 0 {
		if diff < 60 {
			zhTime = strconv.FormatInt(diff, 10) + "秒前"
		} else if diff >= 60 && diff < 60*60 {
			zhTime = strconv.FormatInt(diff/60, 10) + "分钟前"
		} else if diff >= 60*60 && diff < 60*60*24 {
			zhTime = strconv.FormatInt(diff/(60*60), 10) + "小时前"
		} else if diff > 60*60*24 && diff < 60*60*24*365 {
			zhTime = strconv.FormatInt((diff/(60*60*24)), 10) + "天前"
		} else {
			zhTime = "1年前"
		}
	}
	return zhTime
}

func DateTimeFormat(t time.Time) string {
	a := NewArrow(t)
	return a.Format("%Y-%m-%d %H:%M:%S")
}

func DateFormat(t time.Time) string {
	a := NewArrow(t)
	return a.Format("%Y-%m-%d")
}

func MonthFormat(t time.Time) string {
	a := NewArrow(t)
	return a.Format("%Y-%m")
}
