package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

func getIDs(username string) ([]int64, error) {
	t, err := date(*olderThan)
	if err != nil {
		return nil, err
	}

	rsp, err := http.Get(fmt.Sprintf("https://twitter.com/%s/media", username))
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	buf := bytes.NewBuffer(b)
	defer buf.Reset()

	var ids []int64

	for {
		l, err := buf.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if !strings.Contains(l, fmt.Sprintf(`a href="/%s/status/`, username)) {
			continue
		}

		// create parts
		parts := strings.Split(l, " ")
		part := strings.Split(parts[3], "/")[3]
		// get date
		date := strings.Join(parts[8:14], " ")
		date = strings.Split(date, `"`)[1]
		// 3:37 am - 18 Feb 2017
		ti, err := time.Parse("3:04 pm - 2 Jan 2006", date)
		if err != nil {
			return nil, err
		}

		if ti.After(t) {
			continue
		}

		i, _ := strconv.ParseInt(strings.TrimSuffix(part, `"`), 10, 64)
		ids = append(ids, i)
	}

	return ids, nil
}

func media(api *anaconda.TwitterApi) error {
	if len(*username) == 0 {
		return errors.New("require username for media deletion")
	}

	for {
		ids, err := getIDs(*username)
		if err != nil {
			return err
		}

		for _, id := range ids {
			_, err := api.DeleteTweet(id, true)
			if err != nil {
				return fmt.Errorf("error deleting %d: %v", id, err)
			}
			fmt.Println("deleted", id)
		}
	}
}
