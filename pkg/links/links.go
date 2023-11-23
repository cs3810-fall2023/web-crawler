// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 138.
// Package links provides a link-extraction function.
package links

import (
	"fmt"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func Extract(url string) ([]string, error) {
	go Getter(url)
	go Analyser(url)
	return Request(url)
}

func Getter(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	return nil
}

func Analyser(urlRaw string) error {
	_, err := validateURL(urlRaw)
	if err != nil {
		return err
	}
	resp, err := http.Get(urlRaw)
	if err != nil {
		return fmt.Errorf("GET request failed: %s: %v", urlRaw, err)
	}
	_, err = html.Parse(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", urlRaw, err)
	}
	return nil
}

func Request(urlRaw string) ([]string, error) {
	_, err := validateURL(urlRaw)
	if err != nil {
		return nil, err
	}

	resp, err := getResponse(urlRaw)
	if err != nil {
		return nil, err
	}

	links, err := parseHTML(resp, urlRaw)
	if err != nil {
		return nil, err
	}

	return links, nil
}

func getResponse(urlRaw string) (*http.Response, error) {
	resp, err := http.Get(urlRaw)
	if err != nil {
		return nil, fmt.Errorf("GET request failed: %v", err)
	}
	return resp, nil
}

func parseHTML(resp *http.Response, urlRaw string) ([]string, error) {
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", urlRaw, err)
	}

	links := extractLinks(doc, resp)

	return links, nil
}

func extractLinks(doc *html.Node, resp *http.Response) []string {
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links
}

// Validate url
func validateURL(urlRaw string) (*url.URL, error) {
	parsedURL, err := url.ParseRequestURI(urlRaw)
	if err != nil {
		return nil, fmt.Errorf("ERROR: %s is bad", urlRaw)
	}
	return parsedURL, nil
}

// Copied from gopl.io/ch5/outline2.
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
