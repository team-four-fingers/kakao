package local

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/team-four-fingers/kakao/core"
	"github.com/team-four-fingers/kakao/local/keyword"
	"strconv"
)

var (
	ErrNoResult = errors.New("no result")
)

// QueryParams
// query	String	검색을 원하는 질의어	O
// category_group_code	CategoryGroupCode	카테고리 그룹 코드, 카테고리로 결과 필터링을 원하는 경우 사용	X
// x	String	중심 좌표의 X 혹은 경도(longitude) 값
// 특정 지역을 중심으로 검색할 경우 radius와 함께 사용 가능	X
// y	String	중심 좌표의 Y 혹은 위도(latitude) 값
// 특정 지역을 중심으로 검색할 경우 radius와 함께 사용 가능	X
// radius Integer	중심 좌표부터의 반경거리. 특정 지역을 중심으로 검색하려고 할 경우 중심좌표로 쓰일 x,y와 함께 사용
// (단위: 미터(m), 최소: 0, 최대: 20000)

type Client interface {
	SearchByKeyword(query, categoryGroupCode string, x, y float64, radius int) (*keyword.Response, error)
}

type defaultClient struct {
	*resty.Client
}

func NewClient() (Client, error) {
	restyCli := resty.New()

	if err := core.ConfigureClient(restyCli); err != nil {
		return nil, err
	}

	return &defaultClient{Client: restyCli}, nil
}

const (
	queryKey             = "query"
	categoryGroupCodeKey = "category_group_code"
	xKey                 = "x"
	yKey                 = "y"
	radiusKey            = "radius"
)

func (d *defaultClient) SearchByKeyword(query, categoryGroupCode string, x, y float64, radius int) (
	*keyword.Response,
	error,
) {
	respJSON, err := d.NewRequest().
		SetQueryParams(map[string]string{
			queryKey:             query,
			categoryGroupCodeKey: categoryGroupCode,
			xKey:                 strconv.FormatFloat(x, 'f', -1, 64),
			yKey:                 strconv.FormatFloat(y, 'f', -1, 64),
			radiusKey:            strconv.Itoa(radius),
		}).
		Get(keyword.Path)
	if err != nil {
		return nil, err
	}

	var resp *keyword.Response
	if err := json.Unmarshal(respJSON.Body(), &resp); err != nil {
		if jsonErr, ok := err.(*json.SyntaxError); ok {
			problemPart := string(respJSON.Body())[jsonErr.Offset-10 : jsonErr.Offset+10]
			err = fmt.Errorf("%w ~ error near '%s' (offset %d)", err, problemPart, jsonErr.Offset)
		}
		return nil, err
	}

	if len(resp.Documents) == 0 {
		return nil, ErrNoResult
	}

	return resp, nil
}
