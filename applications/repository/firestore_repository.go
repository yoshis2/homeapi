package repository

import (
	"homeapi/domain"

	"cloud.google.com/go/firestore"
)

type FirestoreRepository interface {
	List(*firestore.Client) *firestore.DocumentIterator
	Insert(*firestore.Client, *domain.FirestoreConnect) (*firestore.WriteResult, error)
}
