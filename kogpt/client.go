package kogpt

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/team-four-fingers/kakao/core"
	"github.com/team-four-fingers/kakao/kogpt/generation"
)

const (
	baseURL = "https://api.kakaobrain.com"
)

type Client interface {
	Generate(req *generation.Request) (
		*generation.Response,
		error,
	)
}

type defaultClient struct {
	*resty.Client
}

func NewClient() (Client, error) {
	restyCli := resty.New()

	if err := core.ConfigureClient(restyCli, core.WithBaseURL(baseURL)); err != nil {
		return nil, err
	}

	return &defaultClient{Client: restyCli}, nil
}

func (d *defaultClient) Generate(req *generation.Request) (*generation.Response, error) {
	respJSON, err := d.NewRequest().
		SetBody(req).
		Post(generation.Path)
	if err != nil {
		return nil, err
	}

	var resp *generation.Response
	if err := json.Unmarshal(respJSON.Body(), &resp); err != nil {
		if jsonErr, ok := err.(*json.SyntaxError); ok {
			problemPart := string(respJSON.Body())[jsonErr.Offset-10 : jsonErr.Offset+10]
			err = fmt.Errorf("%w ~ error near '%s' (offset %d)", err, problemPart, jsonErr.Offset)
		}
		return nil, err
	}

	return resp, nil
}
