package core

import "errors"

var globalOptions = struct {
	restAPIKey string
}{
	restAPIKey: "",
}

func InitializeSDK(restAPIKey string) error {
	if restAPIKey == "" {
		return errors.New("REST API Key is empty")
	}

	globalOptions.restAPIKey = restAPIKey
	return nil
}
