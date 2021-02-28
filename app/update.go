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

	//データを更新する
	_, updateError := client.Collection("users").Doc("user2").Set(ctx, map[string]interface{} {
		"first": "editExample.first",
	}, firestore.MergeAll)
	if updateError != nil{
		log.Fatalf("Failed editing alovelace: %s", err)
	}

	//切断
	defer client.Close()
}
