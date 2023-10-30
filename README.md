# CS 3810 Web Crawler

This project implements a simple web crawler that retrieves the HTML content of a given URL and extracts all the links found in the page. The crawler then recursively visits each of the extracted links and repeats the process until a maximum depth is reached or no more links are found.

Usage:
> `go run findlinks.go <url> <depth>`

Arguments:
> url (str): The URL to start crawling from.
> depth (int): The maximum depth to crawl to.

Example:
> `go run findlinks.go https://www.example.com 2`

