package local

import (
	"fmt"
	"github.com/team-four-fingers/kakao/core"
	"os"
)

func ExampleNewClient() {
	kakaoRESTAPIKey := os.Getenv("KAKAO_REST_API_KEY")

	cli := NewClient(core.NewClient(core.WithRestAPIKey(kakaoRESTAPIKey)))

	resp, err := cli.SearchByKeyword("카페", CategoryGroupCode.Cafe, 127.110306812433, 37.394245407468, 10000)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println("response received:", resp.Documents[0].PlaceName)
	// Output: response received: 세시셀라 아브뉴프랑점
}
