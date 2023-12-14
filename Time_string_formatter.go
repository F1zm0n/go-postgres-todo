package main

import (
	"fmt"
	"strings"
	"time"
)

func parseTime(Time string) string {
	parsedTime, err := time.Parse("2006-01-02T15:04:05.999999Z", Time)
	if err != nil {
		fmt.Println("error parsing time", err)
		return ""
	}
	bm := parsedTime.Format("2006-01-02-15:04")
	lastDashIndex := strings.LastIndex(bm, "-")
	newTimeStr := strings.Join([]string{bm[:lastDashIndex], " ", bm[lastDashIndex+1:]}, "")
	return newTimeStr
}
