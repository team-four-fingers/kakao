package core

import (
	"errors"
	"github.com/go-resty/resty/v2"
)

type ClientOptions struct {
	baseURL string
}

type ClientOption func(*ClientOptions)

func WithBaseURL(baseURL string) ClientOption {
	return func(o *ClientOptions) {
		o.baseURL = baseURL
	}
}

func ConfigureClient(restyCli *resty.Client, options ...ClientOption) error {
	o := &ClientOptions{
		baseURL: defaultKakaoApiHost,
	}
	for _, option := range options {
		option(o)
	}

	if globalOptions.restAPIKey == "" {
		return errors.New("REST API key is not set")
	}

	restyCli.SetAuthScheme(kakaoAuthScheme)
	restyCli.SetAuthToken(globalOptions.restAPIKey)
	restyCli.SetBaseURL(o.baseURL)

	return nil
}
