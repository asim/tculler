package main

import (
	cs "encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

func csv(api *anaconda.TwitterApi) error {
	t, err := date(*olderThan)
	if err != nil {
		return err
	}

	f, err := os.Open(*csvFile)
	if err != nil {
		return err
	}
	defer f.Close()
	r := cs.NewReader(f)

	// skip header
	r.Read()

	for {
		rec, err := r.Read()
		if err != nil {
			return err
		}

		id, err := strconv.ParseInt(rec[0], 10, 64)
		if err != nil {
			return err
		}

		ti, err := time.Parse("2006-01-02 15:04:05 -0700", rec[3])
		if err != nil {
			return err
		}

		// skip if newer than t
		if ti.After(t) {
			continue
		}

		_, err = api.DeleteTweet(id, true)
		if err != nil {
			continue
		}
		fmt.Println("deleted", id)
	}
}
