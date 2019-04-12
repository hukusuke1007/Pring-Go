package model

import (
	ballcap "github.com/hukusuke1007/ballcap-go"

	firestore "cloud.google.com/go/firestore"
)

// TweetLikeCollectionPath ...
const TweetLikeCollectionPath = "tweetlike"

// TweetLike ...
type TweetLike struct {
	*ballcap.Base
	Data TweetLikeData
}

// TweetLikeData ...
type TweetLikeData struct {
	UserUID  string `json:"userUid,omitempty"`
	TweetUID string `json:"tweetUid,omitempty"`
}

// NewTweetLike constructor
func NewTweetLike(client *firestore.Client, path *ballcap.Collection, value *ballcap.Value) *TweetLike {
	element := new(TweetLike)
	element.Base = new(ballcap.Base)
	if path == nil {
		path := new(ballcap.Collection)
		path.Path = TweetLikeCollectionPath
	}
	if value != nil {
		element.Init(client, path, value, &element.Data) // CollectionGet
	} else {
		element.Init(client, path, nil, nil) // NewCreate
	}
	return element
}

// Save ...
func (el *TweetLike) Save() error {
	err := el.Base.Save(el.Pack())
	if err != nil {
		return err
	}
	return nil
}

// Update ...
func (el *TweetLike) Update() error {
	err := el.Base.Update(el.Pack())
	if err != nil {
		return err
	}
	return nil
}

// Get ...
func (el *TweetLike) Get() (interface{}, error) {
	err := el.Base.Get(el.ID, &el.Data)
	if err != nil {
		return nil, err
	}
	return el.Data, nil
}

// Pack Marshal対象外の型はここで明示的にmapへ変換する
func (el *TweetLike) Pack() (mapVal map[string]interface{}) {
	val := el.Base.Pack(el.Data)
	return val
}
