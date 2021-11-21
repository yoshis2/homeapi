package repository

import (
	"context"

	"homeapi/domain"

	"cloud.google.com/go/firestore"
)

type FirestoreRepository struct {
}

func (repo *FirestoreRepository) List(ctx context.Context, client *firestore.Client) *firestore.DocumentIterator {
	iter := client.Collection("users").Where("name", "==", "関さん").Documents(ctx)
	return iter
}

func (repo *FirestoreRepository) Insert(ctx context.Context, client *firestore.Client, firestoreConnect *domain.FirestoreConnect) (*firestore.WriteResult, error) {
	documents := client.Collection(firestoreConnect.Collection)
	_, createdAt, err := documents.Add(ctx, map[string]interface{}{
		"name":    firestoreConnect.Name,
		"address": firestoreConnect.Address,
	})
	return createdAt, err
}
