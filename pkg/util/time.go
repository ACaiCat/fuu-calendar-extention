package util

import (
	"errors"
	"strings"
	"time"
)

func ParseExamTime(examTime string) (time.Time, time.Time, error) {
	parts := strings.Split(examTime, "-")
	if len(parts) != 2 {
		return time.Time{}, time.Time{}, errors.New("Invalid examTime: " + examTime)
	}

	startTime, err := time.Parse("15:04", parts[0])
	if err != nil {
		return time.Time{}, time.Time{}, errors.New("Invalid examTime: " + examTime)
	}

	endTime, err := time.Parse("15:04", parts[1])

	if err != nil {
		return time.Time{}, time.Time{}, errors.New("Invalid examTime: " + examTime)
	}

	return startTime, endTime, nil
}
