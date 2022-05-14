package application

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ParseTime(t string) (time.Duration, error) {
	var mins, hours int
	var err error

	parts := strings.SplitN(t, ":", 2)

	switch len(parts) {
	case 1:
		mins, err = strconv.Atoi(parts[0])
		if err != nil {
			return 0, err
		}
	case 2:
		hours, err = strconv.Atoi(parts[0])
		if err != nil {
			return 0, err
		}

		mins, err = strconv.Atoi(parts[1])
		if err != nil {
			return 0, err
		}
	default:
		return 0, fmt.Errorf("invalid time: %s", t)
	}

	if mins > 59 || mins < 0 || hours > 23 || hours < 0 {
		return 0, fmt.Errorf("invalid time: %s", t)
	}

	return time.Duration(hours)*time.Hour + time.Duration(mins)*time.Minute, nil
}
