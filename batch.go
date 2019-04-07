package pring

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
func Batch(client *firestore.Client, bcs []BatchDocument) error {
	ctx := context.Background()
	batch := client.Batch()
	for _, d := range bcs {
		batch.Set(d.DocumentRef, d.Value, firestore.MergeAll)
	}
	_, err := batch.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}
