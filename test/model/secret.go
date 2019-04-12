package model

import (
	ballcap "github.com/hukusuke1007/ballcap-go"

	firestore "cloud.google.com/go/firestore"
)

// SecretCollectionPath ...
const SecretCollectionPath = "secret"

// Secret ...
type Secret struct {
	*ballcap.Base
	Data SecretData
}

// SecretData ...
type SecretData struct {
	UID      string `json:"uid,omitempty"`
	Password string `json:"password,omitempty"`
}

// NewSecret constructor
func NewSecret(client *firestore.Client, path *ballcap.Collection, value *ballcap.Value) *Secret {
	element := new(Secret)
	element.Base = new(ballcap.Base)
	if path == nil {
		path := new(ballcap.Collection)
		path.Path = SecretCollectionPath
	}
	if value != nil {
		element.Init(client, path, value, &element.Data) // CollectionGet
	} else {
		element.Init(client, path, nil, nil) // NewCreate
	}
	return element
}

// Save ...
func (el *Secret) Save() error {
	err := el.Base.Save(el.Pack())
	if err != nil {
		return err
	}
	return nil
}

// Update ...
func (el *Secret) Update() error {
	err := el.Base.Update(el.Pack())
	if err != nil {
		return err
	}
	return nil
}

// Get ...
func (el *Secret) Get() (interface{}, error) {
	err := el.Base.Get(el.ID, &el.Data)
	if err != nil {
		return nil, err
	}
	return el.Data, nil
}

// Pack Marshal対象外の型はここで明示的にmapへ変換する
func (el *Secret) Pack() (mapVal map[string]interface{}) {
	val := el.Base.Pack(el.Data)
	return val
}
