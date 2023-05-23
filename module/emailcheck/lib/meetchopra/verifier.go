package meetchopra

import (
	"encoding/json"
	"net/http"
)

func Verify(email, access_token string) (bool, error) {
	url := "https://verifier.meetchopra.com/verify/" + email + "?token=" + access_token
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	type verfier struct {
		Status bool `json:"status"`
	}

	var result verfier
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return false, err
	}
	return result.Status, err
}
