package main

import (
	"net/url"
	"strings"
)

func NewUrl(ref string) (*url.URL, error) {
	url, err := url.Parse(ref)
	if err != nil {
		return url, err
	}

	if !url.IsAbs() {
		if !strings.Contains(url.Path, "/") {
			url.Path = url.Path + "/" + url.Path
		}
		url.Scheme = "https"
		url.Host = "github.com"
		if url.Path[0] != '/' {
			url.Path = "/" + url.Path
		}
	}

	return url, nil
}
