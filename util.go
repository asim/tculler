package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// 1m | 1h | 1d
func date(age string) (time.Time, error) {
	if len(age) == 0 {
		return time.Now(), nil
	}

	re := regexp.MustCompilePOSIX("([0-9]+)([a-z]+)")

	// blanket invalid age response
	err := func() (time.Time, error) {
		return time.Time{}, fmt.Errorf("%s not valid age. format (1m|1h|1d)", age)
	}

	if !re.MatchString(age) {
		return err()
	}

	parts := re.FindStringSubmatch(age)
	if len(parts) != 3 {
		return err()
	}

	d, er := strconv.Atoi(parts[1])
	if er != nil {
		return err()
	}

	switch parts[2] {
	case "m":
		return time.Now().Add(-time.Minute * time.Duration(d)), nil
	case "h":
		return time.Now().Add(-time.Hour * time.Duration(d)), nil
	case "d":
		return time.Now().Add(-time.Hour * 24 * time.Duration(d)), nil
	}

	return err()
}
