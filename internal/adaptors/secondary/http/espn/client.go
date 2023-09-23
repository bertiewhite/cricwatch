package espn

import "net/http"

type EspnClient struct {
	client *http.Client
}

const (
	baseUrl = "https://site.api.espn.com/apis/site/v2/sports/cricket"
	cacheID = "19435"
)

func NewClient(client *http.Client) *EspnClient {
	return &EspnClient{client}
}
