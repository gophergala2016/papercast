package main

import (
	"github.com/advancedlogic/GoOse"
	"net/url"
)

func extractFromURL(u url.URL) (string, error) {
	g := goose.New()
	article, err := g.ExtractFromURL(u.String())
	if err != nil {
		return "", err
	}

	return article.CleanedText, nil
}
