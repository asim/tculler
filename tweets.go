package main

import (
	"fmt"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
)

func tweets(api *anaconda.TwitterApi) error {
	maxId := ""

	for {
		vals := url.Values{}
		vals.Set("count", "200")
		vals.Set("include_rts", "true")
		vals.Set("exclude_replies", "false")

		if len(maxId) > 0 {
			vals.Set("max_id", maxId)
		}

		tl, err := api.GetUserTimeline(vals)
		if err != nil {
			return err
		}

		for _, t := range tl {
			_, err := api.DeleteTweet(t.Id, true)
			if err != nil {
				return fmt.Errorf("error deleting %d: %v", t.Id, err)
			}
			fmt.Println("deleted", t.Id)
			maxId = fmt.Sprintf("%d", t.Id)
		}
	}
}
