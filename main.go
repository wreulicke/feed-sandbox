package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/feeds"
)

func mainInternal() error {
	f := feeds.Feed{
		Title:       "feed-sandbox",
		Link:        &feeds.Link{Href: "https://github.com/wreulicke/feed-sandbox"},
		Description: "feed-sandbox",
		Author:      &feeds.Author{Name: "wreulicke"},
		// 2024-07-17T23:47:11+09:00
		Created: time.Date(2024, 7, 17, 23, 47, 11, 0, time.FixedZone("JST", 9*60*60)),
	}
	f.Items = []*feeds.Item{
		{
			Title:       "item1",
			Link:        &feeds.Link{Href: "https://github.com/wreulicke/feed-sandbox"},
			Description: "item1",
			Author:      &feeds.Author{Name: "wreulicke"},
			// 2024-07-17T23:47:11+09:00
			Created: time.Date(2024, 7, 17, 23, 47, 11, 0, time.FixedZone("JST", 9*60*60)),
		},
		{
			Title:       "item2",
			Link:        &feeds.Link{Href: "https://github.com/wreulicke/feed-sandbox"},
			Description: "item2",
			Author:      &feeds.Author{Name: "wreulicke"},
			// 2024-07-17T23:48:11+09:00
			Created: time.Date(2024, 7, 17, 23, 48, 11, 0, time.FixedZone("JST", 9*60*60)),
		},
	}
	atom, err := f.ToAtom()
	if err != nil {
		return err
	}
	fmt.Println(atom)

	return nil
}

func main() {
	if err := mainInternal(); err != nil {
		log.Fatal(err)
	}
}
