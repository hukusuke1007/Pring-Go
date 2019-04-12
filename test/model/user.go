package model

import (
	ballcap "github.com/hukusuke1007/ballcap-go"

	firestore "cloud.google.com/go/firestore"
)

// UserCollectionPath ...
const UserCollectionPath = "user"

// User ...
type User struct {
	*ballcap.Base
	Data          UserData // Flat
	SubCollection struct {
		Secret     string
		TweetLikes string
	}
}

// UserData ...
type UserData struct {
	UID      string `json:"uid,omitempty"`
	Name     string `json:"name,omitempty"`
	Age      int64  `json:"age"`
	IsActive bool   `json:"isActive"`
}

// NewUser constructor
func NewUser(client *firestore.Client, path *ballcap.Collection, value *ballcap.Value) *User {
	element := new(User)
	element.Base = new(ballcap.Base)
	if path == nil {
		path = new(ballcap.Collection)
		path.Path = UserCollectionPath
	}
	if value != nil {
		element.Init(client, path, value, &element.Data) // CollectionGet
	} else {
		element.Init(client, path, nil, nil) // NewCreate
	}
	element.SetSubCollection() // SubCollection
	return element
}

// SetSubCollection ...
func (el *User) SetSubCollection() {
	el.SubCollection.Secret = UserCollectionPath + "/" + el.ID + "/" + SecretCollectionPath
	el.SubCollection.TweetLikes = UserCollectionPath + "/" + el.ID + "/" + TweetLikeCollectionPath
}

// Save ...
func (el *User) Save() error {
	err := el.Base.Save(el.Pack())
	if err != nil {
		return err
	}
	return nil
}

// Update ...
func (el *User) Update() error {
	err := el.Base.Update(el.Pack())
	if err != nil {
		return err
	}
	return nil
}

// Get ...
func (el *User) Get() (interface{}, error) {
	err := el.Base.Get(el.ID, &el.Data)
	if err != nil {
		return nil, err
	}
	el.SetSubCollection()
	return el.Data, nil
}

// Pack Marshal対象外の型はここで明示的にmapへ変換する
func (el *User) Pack() (mapVal map[string]interface{}) {
	val := el.Base.Pack(el.Data)
	// for k, v := range val {
	// 	// fmt.Println(k, v)
	// 	if k == "startAt" {
	// 		t, err := time.Parse(time.RFC3339, v.(string))
	// 		if err != nil {
	// 			fmt.Println("time Parse", err)
	// 			break
	// 		}
	// 		mapVal["startAt"] = t
	// 	}
	// 	if k == "endAt" {
	// 		t, err := time.Parse(time.RFC3339, v.(string))
	// 		if err != nil {
	// 			fmt.Println("time Parse", err)
	// 			break
	// 		}
	// 		mapVal["endAt"] = t
	// 	}
	// }
	return val
}
