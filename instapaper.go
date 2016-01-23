package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// InstapaperFeed Unmarshalled Instapaper feed
type InstapaperFeed struct {
	Title string               `xml:"channel>title"`
	Items []InstapaperFeedItem `xml:"channel>item"`
}

// InstapaperFeedItem <item>
type InstapaperFeedItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func (item InstapaperFeedItem) toXML() *RssItem {
	return &RssItem{
		Title:       item.Title,
		Description: item.Description,
		PubDate:     item.PubDate,
		Link:        item.Link,
		GUID:        item.Link,
		Enclosure: &RssEnclosure{
			URL:  proxyLink(item.Link),
			Type: "audio/mpeg",
		},
	}
}

// Instapaper Instapaper
type Instapaper struct {
	userID      string
	hash        string
	feed        *InstapaperFeed
	notModified bool
	etag        string
}

func newInstapaper(userID string, hash string) *Instapaper {
	return &Instapaper{userID: userID, hash: hash}
}

func (i *Instapaper) fetchInstapaperFeed(ifNoneMatch string) error {
	url := fmt.Sprintf("https://www.instapaper.com/rss/%s/%s", i.userID, i.hash)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("If-None-Match", ifNoneMatch)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	i.etag = resp.Header.Get("Etag")

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusNotModified {
		i.notModified = true
		return nil
	}

	i.feed = &InstapaperFeed{}

	err = xml.Unmarshal(data, i.feed)
	if err != nil {
		return err
	}

	return nil
}
