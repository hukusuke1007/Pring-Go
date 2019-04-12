package ballcap

import (
	"context"
	"log"

	firestore "cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// DataSource ...
type DataSource struct {
	CollectionRef *firestore.CollectionRef
	Query         *Query
}

// Value ...
type Value struct {
	ID   string
	Data map[string]interface{}
}

// NewDataSource ...
func NewDataSource(q *Query) *DataSource {
	newDataSource := new(DataSource)
	newDataSource.Query = q
	return newDataSource
}

// Get ...
func (d *DataSource) Get() []Value {
	ctx := context.Background()
	iter := d.Query.Query.Documents(ctx)
	values := []Value{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("error client app: %v\n", err)
		}
		if doc.Exists() {
			value := Value{ID: doc.Ref.ID, Data: doc.Data()}
			values = append(values, value)
		}
	}
	return values
}
