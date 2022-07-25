package macvendorfinder

import (
	"io"
	"net/http"
	"net/url"
	"time"
)

func MacVendorFinder(mac []string) []string {
	var result []string
	for _, ip := range mac {
		query := url.QueryEscape(ip)
		resp, _ := http.Get("https://api.macvendors.com/" + query)
		b, _ := io.ReadAll(resp.Body)
		result = append(result, string(b))
		time.Sleep(time.Second)

	}
	return result
}

func Macfind(mac string) string {
	resp, _ := http.Get("https://api.macvendors.com/" + url.QueryEscape(mac))
	b, _ := io.ReadAll(resp.Body)
	return string(b)
}
