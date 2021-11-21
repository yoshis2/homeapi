package repository

import (
	"context"
	"homeapi/domain"

	"cloud.google.com/go/firestore"
)

type FirestoreRepository interface {
	List(ctx context.Context, client *firestore.Client) *firestore.DocumentIterator
	Insert(ctx context.Context, client *firestore.Client, firestoreConnect *domain.FirestoreConnect) (*firestore.WriteResult, error)
}
