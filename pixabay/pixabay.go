package pixabay

import (
	"github.com/iikira/BaiduPCS-Go/requester"
	"net/url"
)

var (
	// APIKey pixabay api key
	APIKey = "7586785-fa461e94129b4f20ceb83a1a7"
)

// Pixabay pixabay
type Pixabay struct {
	URL    *url.URL
	Client *requester.HTTPClient
	APIKey string

	url string
}

// NewPixabay new pixabay
func NewPixabay() *Pixabay {
	p := &Pixabay{
		URL: &url.URL{
			Scheme: "https",
			Host:   "pixabay.com",
			Path:   "/api",
		},
		Client: requester.NewHTTPClient(),
		APIKey: APIKey,
	}

	p.url = p.URL.String()

	return p
}
