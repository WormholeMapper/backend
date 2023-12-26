package esiclient

import (
	"net/http"

	"github.com/antihax/goesi"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gregjones/httpcache"
	httpmemcache "github.com/gregjones/httpcache/memcache"
)

// CreateCachedTransport initialises an http.Client, with support for ETags and memcached.
func CreateCachedTransport(addresses []string) *http.Client {
	cache := memcache.New(addresses...)
	transport := httpcache.NewTransport(httpmemcache.NewWithClient(cache))
	transport.Transport = &http.Transport{Proxy: http.ProxyFromEnvironment}
	return &http.Client{Transport: transport}
}

func CreateClient(addresses []string, userAgent string) *goesi.APIClient {
	return goesi.NewAPIClient(CreateCachedTransport(addresses),	userAgent);
}