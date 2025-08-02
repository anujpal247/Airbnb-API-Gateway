package util

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func ProxyToServer(targetBaseUrl string, pathPrefix string) http.HandlerFunc {
	target, err := url.Parse(targetBaseUrl)

	if err != nil {
		fmt.Println("Error parsing url", err)
		return nil
	}

	proxy := httputil.NewSingleHostReverseProxy(target)
	originalDireactor := proxy.Director

	proxy.Director = func(r *http.Request) {
		originalDireactor(r)
		originalPath := r.URL.Path

		strippedPath := strings.TrimPrefix(originalPath, pathPrefix)

		r.URL.Host = target.Host
		r.URL.Path = target.Path + strippedPath

		r.Host = target.Host

		if userId, ok := r.Context().Value("userId").(string); ok {
			r.Header.Set("X-User-ID", userId)
		}
	}
	return proxy.ServeHTTP
}
