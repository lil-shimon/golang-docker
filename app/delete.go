package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/api/option"
	"log"

	firebase "firebase.google.com/go"
)

func main() {
	//初期化
	ctx := context.Background()
	sa := option.WithCredentialsFile("path/to/serviceAccount.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	//データを削除する
	_, deleteError := client.Collection("users").Doc("user2").Update(ctx, []firestore.Update {
		{
			Path: "last",
			Value: firestore.Delete,
		},
	})
	if deleteError != nil{
		log.Fatalf("Failed editing alovelace: %s", err)
	}

	//切断
	defer client.Close()
}
