package ballcap

import (
	"context"

	firestore "cloud.google.com/go/firestore"
)

// BatchDocument ...
type BatchDocument struct {
	DocumentRef *firestore.DocumentRef
	Value       map[string]interface{}
}

// Batch ...
type Batch struct {
	Client *firestore.Client
}

// NewBatch constructor
func NewBatch(client *firestore.Client) *Batch {
	element := new(Batch)
	element.Client = client
	return element
}

// Save ...
func (b *Batch) Save(bc BatchDocument) {
	batch := b.Client.Batch()
	batch.Set(bc.DocumentRef, bc.Value, firestore.MergeAll)
}

// Update ...
func (b *Batch) Update(bc BatchDocument) {
	batch := b.Client.Batch()
	batch.Set(bc.DocumentRef, bc.Value, firestore.MergeAll)
}

// Delete ...
func (b *Batch) Delete(bc BatchDocument) {
	batch := b.Client.Batch()
	batch.Delete(bc.DocumentRef)
}

// Commit ...
func (b *Batch) Commit() error {
	ctx := context.Background()
	batch := b.Client.Batch()
	_, err := batch.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}
