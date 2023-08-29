package core

import "github.com/go-resty/resty/v2"

type Client struct {
	*resty.Client
}

type ClientOptions struct {
	restAPIKey string
}

type ClientOption func(*ClientOptions)

func NewClient(options ...ClientOption) *Client {
	o := &ClientOptions{}
	for _, option := range options {
		option(o)
	}

	restyCli := resty.New()

	restyCli.SetAuthScheme(kakaoAuthScheme)
	restyCli.SetAuthToken(o.restAPIKey)
	restyCli.SetBaseURL(defaultKakaoApiHost)

	return &Client{
		restyCli,
	}
}
