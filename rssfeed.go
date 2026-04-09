package main

import "html"

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func (f RSSFeed) decodeHTML() {
	f.Channel.Title = html.UnescapeString(f.Channel.Title)
	f.Channel.Description = html.UnescapeString(f.Channel.Description)

	for _, item := range f.Channel.Item {
		item.decodeHTML()
	}
}

func (i RSSItem) decodeHTML() {
	i.Title = html.UnescapeString(i.Title)
	i.Description = html.UnescapeString(i.Title)
}
