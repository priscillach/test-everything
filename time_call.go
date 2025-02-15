package main

import (
	"fmt"
	"strconv"
	"time"
)

// 实现如下的calc函数
//
// 解题思路：本质是如下周期性曲线的处理，思考如何对题目进行数据结构的简化和抽象，不要陷入到时间格式的处理，这个不是重点
// ----now-----start-------end----------
// ------------start--now--end----------
// ------------start-------end---now----
// ------------end-------start---now----
// -----now----end-------start----------
// ------------end--now--start----------
func calc(input string, start, end string) (startExpect string, endExpect string) {
	t, err := time.Parse("2006-01-02 15:04:05", input)
	if err != nil {
		return "-1", "-1"
	}
	startWeekday, _ := strconv.Atoi(start[9:])
	endWeekday, _ := strconv.Atoi(end[9:])
	if endWeekday < startWeekday {
		endWeekday += 7
	}

	startTime, err := time.Parse("15:04:05", start[:8])
	if err != nil {
		return "-1", "-1"
	}
	endTime, err := time.Parse("15:04:05", end[:8])
	if err != nil {
		return "-1", "-1"
	}

	cur := int(t.Weekday())*24*60*60 + t.Hour()*60*60 + t.Minute()*60 + t.Second()

	startTs := int(startWeekday)*24*60*60 + startTime.Hour()*60*60 + startTime.Minute()*60 + startTime.Second()
	endTs := int(endWeekday)*24*60*60 + endTime.Hour()*60*60 + endTime.Minute()*60 + endTime.Second()
	if cur <= endTs {
		return t.Add(time.Second * time.Duration(startTs-cur)).Format("2006-01-02 15:04:05"), t.Add(time.Second * time.Duration(endTs-cur)).Format("2006-01-02 15:04:05")
	} else {
		return t.Add(time.Second * time.Duration(startTs-cur+7*24*60*60)).Format("2006-01-02 15:04:05"), t.Add(time.Second * time.Duration(endTs-cur+7*24*60*60)).Format("2006-01-02 15:04:05")
	}
	return "-1", "-1"
}

func main() {
	// 常用函数
	// 字符串转时间：     t, err := time.Parse("2006-01-02 15:04:05", input)
	// 时间Add：          func (t Time) Add(d Duration) Time
	// 时间获取星期信息： func (t Time) Weekday() Weekday      样例：int(t.Weedday()) （sunday=0）
	// 时间转string：     func (t Time) Format(layout string) string 样例：time.Format("2006-01-02 15:04:05")
	fmt.Println("Hello, World!")

	check := func(input string, start, end string, startExpect, endExpect string) {
		startActual, endActual := calc(input, start, end)
		if startActual != startExpect || endActual != endExpect {
			fmt.Printf("[FAILED] input:%v start:%v end:%v expect:[%v - %v] actual:[%v - %v]\n", input, start, end, startExpect, endExpect, startActual, endActual)
		} else {
			fmt.Printf("[SUCCESS] input:%v start:%v end:%v", input, start, end)
		}
	}

	// 检查如下测试用例，最后两个参数是期望返回的时间范围
	check("2024-09-27 16:40:00", "20:00:00 1", "02:00:00 3", "2024-09-30 20:00:00", "2024-10-02 02:00:00")
	check("2024-09-24 16:40:00", "20:00:00 1", "02:00:00 3", "2024-09-23 20:00:00", "2024-09-25 02:00:00")
	check("2024-09-24 16:40:00", "20:00:00 3", "02:00:00 1", "2024-09-25 20:00:00", "2024-09-30 02:00:00")
	check("2024-09-27 16:40:00", "20:00:00 3", "02:00:00 1", "2024-09-25 20:00:00", "2024-09-30 02:00:00")
	check("2024-09-27 16:40:00", "20:00:00 7", "02:00:00 1", "2024-09-29 20:00:00", "2024-09-30 02:00:00")
}
