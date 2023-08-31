package local

import (
	"fmt"
	"github.com/team-four-fingers/kakao/core"
	"os"
)

func ExampleNewClient() {
	kakaoRESTAPIKey := os.Getenv("KAKAO_REST_API_KEY")

	cli := NewClient(core.NewClient(core.WithRestAPIKey(kakaoRESTAPIKey)))

	resp, err := cli.SearchByKeyword("카페", CategoryGroupCode.Cafe, 127.108622, 37.401219, 20000)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println("response received:", resp.Documents[0].PlaceName)
	// Output: response received: 알레그리아 판교테크노밸리점
}
