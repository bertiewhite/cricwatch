package espn

import "net/http"

type Client struct {
	client *http.Client
}

const (
	baseUrl = "https://site.api.espn.com/apis/site/v2/sports/cricket"
	cacheID = "19434"
)

func NewClient(client *http.Client) *Client {
	return &Client{client}
}
