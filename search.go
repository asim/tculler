package main

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

func search(api *anaconda.TwitterApi) error {
	td, err := date(*olderThan)
	if err != nil {
		return err
	}

	if len(*username) == 0 {
		return errors.New("require username for search deletion")
	}

	maxId := ""

	fmt.Println("searching for", fmt.Sprintf("@%s", *username))

	for {
		vals := url.Values{}
		vals.Set("count", "100")

		if len(maxId) > 0 {
			vals.Set("max_id", maxId)
		}

		sr, err := api.GetSearch(fmt.Sprintf("@%s", *username), vals)
		if err != nil {
			return err
		}

		if len(maxId) > 0 {
			fmt.Println("at id", maxId)
		}

		for _, t := range sr.Statuses {
			maxId = fmt.Sprintf("%d", t.Id)

			if t.User.ScreenName != *username {
				continue
			}

			ti, err := time.Parse(time.RubyDate, t.CreatedAt)
			if err != nil {
				return err
			}

			if ti.After(td) {
				continue
			}

			_, err = api.DeleteTweet(t.Id, true)
			if err != nil {
				return fmt.Errorf("error deleting %d: %v", t.Id, err)
			}

			fmt.Println("deleted", t.Id)
		}
	}
}
