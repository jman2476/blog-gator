package main

import (
	"context"
	"encoding/xml"
	"io"
	"net/http"
)

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, nil
	}

	req.Header.Set("User-Agent", "gator")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return &RSSFeed{}, nil
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return &RSSFeed{}, nil
	}

	var feed RSSFeed
	err = xml.Unmarshal(data, &feed)
	if err != nil {
		return &RSSFeed{}, nil
	}

	feed.decodeHTML()

	return &feed, nil
}
