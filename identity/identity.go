package identity

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type LoginResponse struct {
	Kind         string `json:"kind"`
	LocalID      string `json:"localId"`
	Email        string `json:"email"`
	DisplayName  string `json:"displayName"`
	IDToken      string `json:"idToken"`
	Registered   bool   `json:"registered"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
}

var client = resty.New().R()

// Send Request login to identity-toolkit api
func SendRequestLogin(email, password string) (*LoginResponse, error) {
	result := new(LoginResponse)
	res, err := client.SetHeader("content-type", "application/json").
		SetBody(map[string]interface{}{
			"email":    email,
			"password": password,
		}).
		SetResult(result).
		Post("https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=AIzaSyCUiotIte0FYHOK3Zo9h44Qwyo1OtVFlfI")

	if err != nil {
		return nil, err
	}
	fmt.Printf("SendRequestLogin: %+v\n", res)
	return result, nil
}

// Send Request Exhcange Token to identity-toolkit api
func ExchangeCustomIdToken(token string) (*LoginResponse, error) {
	result := new(LoginResponse)
	res, err := client.SetHeader("content-type", "application/json").
		SetBody(map[string]interface{}{
			"token":             token,
			"returnSecureToken": true,
		}).
		SetResult(result).
		Post("https://identitytoolkit.googleapis.com/v1/accounts:signInWithCustomToken?key=AIzaSyCUiotIte0FYHOK3Zo9h44Qwyo1OtVFlfI")

	if err != nil {
		return nil, err
	}
	fmt.Printf("Exchange Auth Custom Token: %+v\n", res)
	return result, nil
}

// Send Request Exhcange Token to identity-toolkit api
func RefreshToken(refreshToken string) (*LoginResponse, error) {
	result := new(LoginResponse)
	res, err := client.SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(map[string]string{
			"grant_type":    "refresh_token",
			"refresh_token": refreshToken,
		}).
		SetResult(result).
		Post("https://securetoken.googleapis.com/v1/token?key=AIzaSyCUiotIte0FYHOK3Zo9h44Qwyo1OtVFlfI")

	if err != nil {
		return nil, err
	}
	fmt.Printf("Request RefreshToken Token: %+v\n", res)
	return result, nil
}
