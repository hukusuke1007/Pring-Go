package pring

import (
	firestore "cloud.google.com/go/firestore"
)

// Query ...
type Query struct {
	IsReference bool
	Reference   *firestore.CollectionRef
	Query       firestore.Query
	Element     interface{}
}

// NewQuery ...
func NewQuery(reference *firestore.CollectionRef, query firestore.Query, IsReference bool) *Query {
	newQuery := new(Query)
	newQuery.Reference = reference
	newQuery.Query = query
	newQuery.IsReference = IsReference
	return newQuery
}

// Where ...
func (q *Query) Where(path, op string, value interface{}) *Query {
	query := NewQuery(q.Reference, q.Query, q.IsReference)
	q.Query = query.Query.Where(path, op, value)
	return q
}

// OrderBy ...
func (q *Query) OrderBy(path string, dir firestore.Direction) *Query {
	query := NewQuery(q.Reference, q.Query, q.IsReference)
	q.Query = query.Query.OrderBy(path, dir)
	return q
}

// Limit ...
func (q *Query) Limit(n int) *Query {
	query := NewQuery(q.Reference, q.Query, q.IsReference)
	q.Query = query.Query.Limit(n)
	return q
}

// StartAt ...
func (q *Query) StartAt(docSnapshotOrFieldValues ...interface{}) *Query {
	query := NewQuery(q.Reference, q.Query, q.IsReference)
	q.Query = query.Query.StartAt(docSnapshotOrFieldValues)
	return q
}

// StartAfter ...
func (q *Query) StartAfter(docSnapshotOrFieldValues ...interface{}) *Query {
	query := NewQuery(q.Reference, q.Query, q.IsReference)
	q.Query = query.Query.StartAfter(docSnapshotOrFieldValues)
	return q
}

// EndAt ...
func (q *Query) EndAt(docSnapshotOrFieldValues ...interface{}) *Query {
	query := NewQuery(q.Reference, q.Query, q.IsReference)
	q.Query = query.Query.EndAt(docSnapshotOrFieldValues)
	return q
}

// EndBefore ...
func (q *Query) EndBefore(docSnapshotOrFieldValues ...interface{}) *Query {
	query := NewQuery(q.Reference, q.Query, q.IsReference)
	q.Query = query.Query.EndBefore(docSnapshotOrFieldValues)
	return q
}

// DataSource ...
func (q *Query) DataSource() *DataSource {
	return NewDataSource(q)
}
