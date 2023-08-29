package core

import "fmt"

func MakeAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf("%s %s", kakaoAuthScheme, accessToken)
}
