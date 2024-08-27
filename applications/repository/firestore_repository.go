package repository

import (
	"context"
	"homeapi/domain"
	"homeapi/interfaces/repository"

	"cloud.google.com/go/firestore"
)

//go:generate mockgen -package mock -source $GOFILE -destination mock/$GOFILE
func NewFirestoreRepository(client *firestore.Client) repository.FirestoreRepository {
	return repository.FirestoreRepository{
		Client: client,
	}
}

type FirestoreRepository interface {
	List(ctx context.Context, client *firestore.Client) *firestore.DocumentIterator
	Insert(ctx context.Context, client *firestore.Client, firestoreConnect *domain.FirestoreConnect) (*firestore.WriteResult, error)
}
