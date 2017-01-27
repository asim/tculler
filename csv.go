package main

import (
	cs "encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
)

func csv(api *anaconda.TwitterApi) error {
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

		_, err = api.DeleteTweet(id, true)
		if err != nil {
			continue
		}
		fmt.Println("deleted", id)
	}
}
