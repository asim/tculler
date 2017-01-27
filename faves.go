package main

import (
	"fmt"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
)

func faves(api *anaconda.TwitterApi) error {
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

		for _, t := range tl {
			_, err := api.Unfavorite(t.Id)
			if err != nil {
				return fmt.Errorf("error unfaving %d: %v", t.Id, err)
			}
			fmt.Println("unfaving", t.Id)
			maxId = fmt.Sprintf("%d", t.Id)
		}
	}
}
