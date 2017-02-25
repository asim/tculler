package main

import (
	"fmt"
	"net/url"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

func faves(api *anaconda.TwitterApi) error {
	t, err := date(*olderThan)
	if err != nil {
		return err
	}

	vals := url.Values{}
	vals.Set("count", "200")

	maxId := ""

	for {
		vals := url.Values{}
		vals.Set("count", "200")

		if len(maxId) > 0 {
			vals.Set("max_id", maxId)
		}

		tl, err := api.GetFavorites(vals)
		if err != nil {
			return err
		}

		for _, tw := range tl {
			ti, err := time.Parse(time.RubyDate, tw.CreatedAt)
			if err != nil {
				return err
			}

			if ti.After(t) {
				continue
			}

			_, err = api.Unfavorite(tw.Id)
			if err != nil {
				return fmt.Errorf("error unfaving %d: %v", tw.Id, err)
			}
			fmt.Println("unfaving", tw.Id)
			maxId = fmt.Sprintf("%d", tw.Id)
		}
	}
}
