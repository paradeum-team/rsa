package utils

import (
	"strconv"
	"time"
)

func GetCurrentTimeUnix() string {
	//当前时间戳
	t1 := time.Now().Unix()
	timeContent := strconv.FormatInt(t1,10)
	return timeContent
}
