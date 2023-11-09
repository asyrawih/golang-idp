package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"github.com/go-resty/resty/v2"
)

func main() {
	ctx := context.Background()

	lr, err := SendRequestLogin("hasyrawi@gmail.com", "awiroot123")
	if err != nil {
		log.Fatal(err)
	}

	app, err := firebase.NewApp(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	c, err := app.Auth(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}

	s, err := c.CustomTokenWithClaims(ctx, lr.LocalID, map[string]interface{}{
		"user_id":     lr.LocalID,
		"email":       lr.Email,
		"device_info": "some device info in here",
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	token, err := ExchangeCustomIdToken(s)
	if err != nil {
		log.Fatal(err.Error())
	}

	// verify id token with revocetion to tell is not has revoked
	t, err := c.VerifyIDTokenAndCheckRevoked(ctx, token.IDToken)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("t: %+v\n", t)

}

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
	_, err := client.SetHeader("content-type", "application/json").
		SetBody(map[string]interface{}{
			"email":    email,
			"password": password,
		}).
		SetResult(result).
		Post("https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=AIzaSyCUiotIte0FYHOK3Zo9h44Qwyo1OtVFlfI")

	if err != nil {
		return nil, err
	}
	return result, nil
}

// Send Request login to identity-toolkit api
func ExchangeCustomIdToken(token string) (*LoginResponse, error) {
	result := new(LoginResponse)
	_, err := client.SetHeader("content-type", "application/json").
		SetBody(map[string]interface{}{
			"token":             token,
			"returnSecureToken": true,
		}).
		SetResult(result).
		Post("https://identitytoolkit.googleapis.com/v1/accounts:signInWithCustomToken?key=AIzaSyCUiotIte0FYHOK3Zo9h44Qwyo1OtVFlfI")

	if err != nil {
		return nil, err
	}
	return result, nil
}
