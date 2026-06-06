package util

import (
	"fmt"
	"strings"
	"time"
)

func ParseExamTime(examData, examTime string) (time.Time, time.Time, error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")

	parts := strings.Split(examTime, "-")
	if len(parts) != 2 {
		return time.Time{}, time.Time{}, fmt.Errorf("parse exam time: %s %s", examTime, examData)
	}
	startTime, err := time.ParseInLocation("2006年01月02日15:04", examData+parts[0], loc)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("parse exam time: %w %s %s", err, examTime, examData)
	}

	endTime, err := time.ParseInLocation("2006年01月02日15:04", examData+parts[1], loc)

	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("parse exam time: %w %s %s", err, examTime, examData)
	}

	return startTime, endTime, nil
}
