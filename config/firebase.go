package config

import (
	"context"
	"fmt"
	"path/filepath"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var FB *auth.Client

func SetupFirebase() *auth.Client {
	serviceAccountKeyFilePath, err := filepath.Abs("./serviceAccountKey.json")
	if err != nil {
		fmt.Println("Unable to load serviceAccountKey.json file:")
	}

	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Println("Firebase Setup Error:", err)
	}

	firebaseAuth, err := app.Auth(context.Background())
	if err != nil {
		fmt.Println("Firebase Setup Error:", err)
	}

	return firebaseAuth
}
