package ballcap

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	firestore "cloud.google.com/go/firestore"
)

// Version ...
var Version = "1"

// Base ...
type Base struct {
	ID            string
	Path          string
	Client        *firestore.Client
	CollectionRef *firestore.CollectionRef
	TimeAt        TimeAt
	value         map[string]interface{}
	IsSaved       bool
}

// TimeAt ...
type TimeAt struct {
	CreatedAt time.Time `json:"createdAt,string"`
	UpdatedAt time.Time `json:"updatedAt,string"`
}

// Collection ...
type Collection struct {
	ID   string
	Path string // version/1 以降のパス
}

// NewCollection ...
func NewCollection(client *firestore.Client, path string) *Query {
	reference := GetReference(client, path)
	return NewQuery(reference, reference.Query, true)
}

// GetReference ...
func GetReference(client *firestore.Client, path string) *firestore.CollectionRef {
	return client.Collection("version/" + Version + "/" + path)
}

// Init ...
func (b *Base) Init(client *firestore.Client, collection *Collection, value *Value, data interface{}) {
	b.Path = "version/" + Version + "/" + collection.Path
	b.Client = client
	b.CollectionRef = b.Client.Collection(b.Path)
	b.IsSaved = false
	if value != nil {
		b.IsSaved = true
		b.ID = value.ID
		b.MapToStruct(value.Data, &data)
	} else {
		if len(collection.ID) == 0 {
			b.ID = b.CollectionRef.NewDoc().ID
		} else {
			b.ID = collection.ID
		}
	}
	fmt.Println("Collection", collection)
}

// Save ...
func (b *Base) Save(value map[string]interface{}) error {
	ctx := context.Background()
	r, err := b.CollectionRef.Doc(b.ID).Set(ctx, value, firestore.MergeAll)
	if err != nil {
		return err
	}
	b.value = value
	b.IsSaved = true
	fmt.Println("Save", r)
	return nil
}

// Update ...
func (b *Base) Update(value map[string]interface{}) error {
	ctx := context.Background()
	r, err := b.CollectionRef.Doc(b.ID).Set(ctx, value, firestore.MergeAll)
	if err != nil {
		return err
	}
	b.value = value
	b.IsSaved = true
	fmt.Println("Update", r)
	return nil
}

// Get ...
func (b *Base) Get(ID string, data interface{}) error {
	ctx := context.Background()
	dsnap, err := b.CollectionRef.Doc(ID).Get(ctx)
	if err != nil {
		return err
	}
	if dsnap.Exists() {
		b.MapToStruct(dsnap.Data(), data)
		b.ID = ID
	}
	b.value = dsnap.Data()
	b.IsSaved = true
	fmt.Println("Get", data)
	return nil
}

// Delete ...
func (b *Base) Delete() error {
	ctx := context.Background()
	r, err := b.CollectionRef.Doc(b.ID).Delete(ctx)
	if err != nil {
		return err
	}
	b.value = nil
	fmt.Println("Delete", r)
	return nil
}

// Pack ...
func (b *Base) Pack(data interface{}) (mapVal map[string]interface{}) {
	mapVal = make(map[string]interface{})
	jsonVal, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Marshal err", err)
		return
	}

	err = json.Unmarshal(jsonVal, &mapVal)
	if err != nil {
		fmt.Println("Unmarshal err", err)
		return
	}

	if b.IsSaved == true {
		timeAt := time.Now()
		mapVal["updatedAt"] = timeAt
	} else {
		timeAt := time.Now()
		mapVal["createdAt"] = timeAt
		mapVal["updatedAt"] = timeAt
	}
	return mapVal
}

// GetBatchDocument ...
func (b *Base) GetBatchDocument(value map[string]interface{}) BatchDocument {
	return BatchDocument{DocumentRef: b.CollectionRef.Doc(b.ID), Value: value}
}

// MapToStruct ...
func (b *Base) MapToStruct(mapVal map[string]interface{}, data interface{}) error {
	jsonVal, err := json.Marshal(mapVal)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonVal, data)
	if err != nil {
		return err
	}
	if val, ok := mapVal["createdAt"]; ok == true {
		b.TimeAt.CreatedAt = val.(time.Time)
	}
	if val, ok := mapVal["updatedAt"]; ok == true {
		b.TimeAt.UpdatedAt = val.(time.Time)
	}
	return nil
}
