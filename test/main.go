package main

import (
	"context"
	"fmt"
	"log"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	pring "github.com/hukusuke1007/pring-go"
	model "github.com/hukusuke1007/pring-go/test/model"
	"google.golang.org/api/option"
)

func main() {
	firebaseDocumentSample()
	// firebaseCollectionSample()
}

func firebaseInit() *firestore.Client {
	opt := option.WithCredentialsFile("test/secretKey/serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
		return nil
	}
	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("error client app: %v\n", err)
		return nil
	}
	return client
}

func firebaseDocumentSample() {

	client := firebaseInit()
	var ID string
	// Save
	{
		user := model.NewUser(client, nil, nil)
		user.Data.Name = "pring-go"
		user.Data.Age = 100
		err := user.Save()
		if err != nil {
			fmt.Println("err", err)
		}
		ID = user.ID
	}

	// Get
	{
		user := model.NewUser(client, &pring.Collection{ID: ID, Path: model.UserCollectionPath}, nil)
		_, err := user.Get()
		if err != nil {
			fmt.Println("err", err)
		}
		fmt.Println("result", user.Data, user.TimeAt)
	}

	// Delete
	{
		// user := model.NewUser(client, &pring.Collection{ID: ID, Path: model.UserCollectionPath}, nil)
		// err := user.Delete()
		// if err != nil {
		// 	fmt.Println("err", err)
		// }
		// fmt.Println("result", user.Data)
	}
}

func firebaseCollectionSample() {

	firebaseCollectionSet()
	// firebaseCollectionGet()
}

func firebaseCollectionSet() {
	// path := "user"
	client := firebaseInit()

	// User Secret
	{
		user := model.NewUser(client, nil, nil)
		user.Data.Name = "aiueo"
		user.Data.Age = 1
		user.Save()

		secret := model.NewSecret(client, &pring.Collection{Path: user.SubCollection.Secret}, nil)
		secret.Data.Password = "password"
		secret.Save()

		user.Data.Age = 200
		user.Update()

		secret.Data.Password = "popopop"
		secret.Update()

		// 	{
		// 		d, e := user.Get()
		// 		if e != nil {
		// 			fmt.Println(e)
		// 			return
		// 		}
		// 		fmt.Println("USER", d)
		// 	}

		// 	{
		// 		d, e := secret.Get()
		// 		if e != nil {
		// 			fmt.Println(e)
		// 			return
		// 		}
		// 		fmt.Println("SECRET", d)
		// 	}
		// }
	}

	// For Batch
	{
		user := model.NewUser(client, nil, nil)
		user.Data.Name = "testtest"
		user.Data.Age = 20

		secret := model.NewSecret(client, &pring.Collection{Path: user.SubCollection.Secret}, nil)
		secret.Data.Password = "passworddes"

		documents := []pring.BatchDocument{}
		documents = append(documents, user.GetBatchDocument(user.Pack()))
		documents = append(documents, secret.GetBatchDocument(secret.Pack()))
		pring.Batch(client, documents)
	}

	// // Data set
	// {
	// 	// Prepare
	// 	ctx := context.Background()
	// 	ref := client.Collection(path)
	// 	users := []pring.User{}
	// 	for _, u := range users {
	// 		data := u.Pack()
	// 		if _, err := ref.Doc(u.Data.UID).Set(ctx, data, firestore.MergeAll); err != nil {
	// 			log.Fatalf("error client app: %v\n", err)
	// 		}
	// 	}
	// }
}

func firebaseCollectionGet() {
	path := model.UserCollectionPath
	collectionPath := &pring.Collection{Path: path}
	client := firebaseInit()

	// Datasource
	{
		users := []model.User{}
		values := pring.NewCollection(client, path).Where("isActive", "==", false).Limit(10).DataSource().Get()
		for _, value := range values {
			user := model.NewUser(client, collectionPath, &value)
			users = append(users, *user)
		}
		for i, value := range users {
			fmt.Println("users", i, value.Path, value.ID, value.Data, value.TimeAt)
		}
	}

	// {
	// 	// Prepare
	// 	ctx := context.Background()
	// 	ref := client.Collection(path)
	// 	// Get
	// 	query := ref.Where("isActive", "==", true)
	// 	iter := query.Documents(ctx)
	// 	for {
	// 		doc, err := iter.Next()
	// 		if err == iterator.Done {
	// 			break
	// 		}
	// 		if err != nil {
	// 			log.Fatalf("error client app: %v\n", err)
	// 		}
	// 		data := doc.Data()
	// 		user := pring.User{&pring.Base{}, pring.UserData{}}
	// 		if err := user.MapToStruct(data, &user.Data); err == nil {
	// 			fmt.Println(user.Data)
	// 		}
	// 	}
	// }
}
