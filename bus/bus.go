package bus

import "net/url"

type bus struct {
	scheme string
	host   string
}

func NewBus(scheme string, host string) *bus {
	return &bus{scheme: scheme, host: host}
}

func (b *bus) createUrl(baseUrl string, params url.Values) string {
	targetUrl := url.URL{
		Scheme:   b.scheme,
		Host:     b.host,
		Path:     baseUrl,
		RawQuery: params.Encode(),
	}
	return targetUrl.String()
}
