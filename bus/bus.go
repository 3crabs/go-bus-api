package bus

import "net/url"

type bus struct {
	scheme string
	host   string
	path   string
}

func NewBus(scheme, host, path string) *bus {
	return &bus{
		scheme: scheme,
		host:   host,
		path:   path,
	}
}

func (b *bus) createUrl(baseUrl string, params url.Values) string {
	targetUrl := url.URL{
		Scheme:   b.scheme,
		Host:     b.host,
		Path:     b.path + baseUrl,
		RawQuery: params.Encode(),
	}
	return targetUrl.String()
}
