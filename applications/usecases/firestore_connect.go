package usecases

import (
	"context"
	"fmt"

	"homeapi/domain"

	"homeapi/infrastructure/firebases"

	"homeapi/applications/logging"
	"homeapi/applications/ports"
	"homeapi/applications/repository"
	"homeapi/applications/util"

	"github.com/go-playground/validator/v10"
)

type FirestoreConnectUsecase struct {
	FirestoreRepository repository.FirestoreRepository
	Firestore           *firebases.Firestore
	Logging             logging.Logging
	Validator           *validator.Validate
}

func (usecase *FirestoreConnectUsecase) List(ctx context.Context) (*[]ports.FirestoreConnectOutputPort, error) {
	client, err := usecase.Firestore.Open()
	if err != nil {
		usecase.Logging.Error(err)
		return nil, err
	}

	iter := usecase.FirestoreRepository.List(ctx, client)
	docs, err := iter.GetAll()
	if err != nil {
		usecase.Logging.Error(err)
		return nil, err
	}

	defer client.Close()

	outputs := make([]ports.FirestoreConnectOutputPort, len(docs))
	for i, doc := range docs {
		if doc.Data()["name"] == nil {
			return nil, fmt.Errorf("BadRequest 名前が存在しないデータがあります。")
		}

		if doc.Data()["address"] == nil {
			return nil, fmt.Errorf("BadRequest 住所が存在しないデータがあります。")
		}

		now, err := util.JapaneseNowTime()
		if err != nil {
			usecase.Logging.Error(err)
			return nil, err
		}

		outputs[i] = ports.FirestoreConnectOutputPort{
			Name:       doc.Data()["name"].(string),
			Address:    doc.Data()["address"].(string),
			Created_at: now,
		}
	}
	return &outputs, nil
}

func (usecase *FirestoreConnectUsecase) Create(ctx context.Context, input *ports.FirestoreConnectInputPort) (*ports.FirestoreConnectOutputPort, error) {
	firestoreConnect := &domain.FirestoreConnect{
		Collection: input.Collection,
		Name:       input.Name,
		Address:    input.Address,
	}

	client, err := usecase.Firestore.Open()
	if err != nil {
		usecase.Logging.Error(err)
		return nil, err
	}

	createdAt, err := usecase.FirestoreRepository.Insert(ctx, client, firestoreConnect)
	if err != nil {
		usecase.Logging.Error(err)
		return nil, err
	}

	defer client.Close()

	output := &ports.FirestoreConnectOutputPort{
		Name:       firestoreConnect.Name,
		Address:    firestoreConnect.Address,
		Created_at: createdAt.UpdateTime,
	}

	return output, nil
}
