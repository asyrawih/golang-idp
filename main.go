package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
)

func main() {
	ctx := context.Background()

	app, err := firebase.NewApp(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	c, err := app.Auth(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}

	ur, err := c.GetUserByEmail(ctx, "hanan@asyrawih.id")

	if err != nil {
		log.Fatal(err)
	}

	token, err := c.CustomTokenWithClaims(ctx, ur.UID, map[string]interface{}{
		"user":    "hanan",
		"user_id": "test",
	})
	if err != nil {
		log.Fatal(err)
	}

	//
	// ❮ curl 'https://identitytoolkit.googleapis.com/v1/accounts:signInWithCustomToken?key=AIzaSyCUiotIte0FYHOK3Zo9h44Qwyo1OtVFlfI' \
	//       -H 'Content-Type: application/json' \
	//       --data-binary '{"token":"custom_token_from_backend_generated","returnSecureToken":true}'

	t, err := c.VerifyIDToken(
		ctx,
		"eyJhbGciOiJSUzI1NiIsImtpZCI6ImQ0OWU0N2ZiZGQ0ZWUyNDE0Nzk2ZDhlMDhjZWY2YjU1ZDA3MDRlNGQiLCJ0eXAiOiJKV1QifQ.eyJ1c2VyIjoiaGFuYW4iLCJ1c2VyX2lkIjoiSDVRQ2hIeXZLZWFrbHFTYVNSaWEweFpUS0k0MiIsImlzcyI6Imh0dHBzOi8vc2VjdXJldG9rZW4uZ29vZ2xlLmNvbS9kYW50ZS02NjYiLCJhdWQiOiJkYW50ZS02NjYiLCJhdXRoX3RpbWUiOjE2OTk1NDc3MjcsInN1YiI6Ikg1UUNoSHl2S2Vha2xxU2FTUmlhMHhaVEtJNDIiLCJpYXQiOjE2OTk1NDc3MjcsImV4cCI6MTY5OTU1MTMyNywiZW1haWwiOiJoYW5hbkBhc3lyYXdpaC5pZCIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwiZmlyZWJhc2UiOnsiaWRlbnRpdGllcyI6eyJlbWFpbCI6WyJoYW5hbkBhc3lyYXdpaC5pZCJdfSwic2lnbl9pbl9wcm92aWRlciI6ImN1c3RvbSJ9fQ.V7S6q3o-1MHOMO0HC7Ya1kOArOQ7KUSbmjJ_2f8EkXDp8ueD1kKN_q3OiQnzOjRl8TUfJqzJoAyyFj4r7kV2FS-VD3OmeQ5CarOH1NrKq6MNGvyBgmIuxPQxj_4igwfX4FUMLymFEyIDWbv9MD50R0bY3w1Qr3BslSbgrG89lLMjcK3WqVdNUYadWgCO8UCQH_WuLIXdf9cNIQ2Frc26yzDka5f2V5PFXI7mW0t7s0rZshXpV2Z-5sk6oDHHsuEbEowS-obrAvlxo9X7Kw41n1oOkYPITPP-TUW49mN05L2yXKPafFJCC3lW6G89hp_tVD1hWpuYxS30iySq8kVBEQ",
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("t.Claims: %v\n", t.Claims["user"])

	fmt.Printf("token: %v\n", token)

}
