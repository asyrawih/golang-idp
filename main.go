package main

import (
	"context"
	"fmt"
	"log"

	"identity/identity"

	firebase "firebase.google.com/go"
)

func main() {
	ctx := context.Background()

	lr, err := identity.SendRequestLogin("hasyrawi@gmail.com", "awiroot123")
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

	fmt.Printf("s: %v\n", s)

	token, err := identity.ExchangeCustomIdToken(s)
	if err != nil {
		log.Fatal(err.Error())
	}

	newtoken, err := identity.RefreshToken(token.RefreshToken)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("newtoken: %v\n", newtoken)

	// verify id token with revocetion to tell is not has revoked
	t, err := c.VerifyIDTokenAndCheckRevoked(ctx, token.IDToken)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("t: %+v\n", t)

}
