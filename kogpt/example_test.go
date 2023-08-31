package kogpt

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/team-four-fingers/kakao/core"
	"github.com/team-four-fingers/kakao/kogpt/generation"
	"os"
	"strings"
)

func ExampleNewClient() {
	kakaoRESTAPIKey := os.Getenv("KAKAO_REST_API_KEY")

	if err := core.InitializeSDK(kakaoRESTAPIKey); err != nil {
		fmt.Println("err:", err)
		return
	}

	cli, err := NewClient()
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	resp, err := cli.Generate(&generation.Request{
		Prompt: `
		"데이트하기 좋은 스파게티 전문점"을 가고 싶은 사람을 위해 카카오맵 키워드 검색을 실행하려고 합니다.

		적절한 키워드들을 ,로 구분하여 전달해주세요.
`,
		//		Prompt: `
		//		주제에 맞는 음식이나 식재료들 json 데이터를 생성하여 파싱하고 싶습니다.
		//        name은 상품의 이름, price는 상품의 가격을 의미합니다. image_url은 ""로 둡니다.
		//
		//		예시:
		//		[
		//			{
		//				"name":     "불고기",
		//				"price":    10000,
		//				"image_url": "",
		//			},
		//		]
		//
		//		"가벼운"이라는 주제에 맞는 음식이나 식재료들을 위와 같은 구조의 json으로 생성해주세요.
		//`,
		MaxTokens:      generation.MaxTokenInSingleConversation,
		TopProbability: lo.ToPtr(float64(0)),
		Temperature:    lo.ToPtr(float64(0)),
	})
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	generatedTexts := lo.Map(resp.Generations, func(item generation.Generation, _ int) string {
		return item.Text
	})

	fmt.Println("response received:", strings.Join(generatedTexts, "\n"))
	// Output: response received:
}
