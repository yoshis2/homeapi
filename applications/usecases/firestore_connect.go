package usecases

import (
	"fmt"

	"homeapi/domain"

	"homeapi/infrastructure/firebases"

	"homeapi/applications/logging"
	"homeapi/applications/ports"
	"homeapi/applications/repository"
	"homeapi/applications/util"
)

type FirestoreConnectUsecase struct {
	FirestoreRepository repository.FirestoreRepository
	Firestore           *firebases.Firestore
	Logging             logging.Logging
}

func (usecase *FirestoreConnectUsecase) List() (*[]ports.FirestoreConnectOutputPort, error) {
	client, err := usecase.Firestore.Open()
	if err != nil {
		usecase.Logging.Error(err)
		return nil, err
	}

	iter := usecase.FirestoreRepository.List(client)
	docs, err := iter.GetAll()
	if err != nil {
		usecase.Logging.Error(err)
		return nil, err
	}

	outputs := make([]ports.FirestoreConnectOutputPort, len(docs))
	for i, doc := range docs {
		if doc.Data()["name"] == nil {
			return nil, fmt.Errorf("BadRequest 名前が存在しないデータがあります。")
		}

		if doc.Data()["address"] == nil {
			return nil, fmt.Errorf("BadRequest 住所が存在しないデータがあります。")
		}

		japaneseTime, err := util.JapaneseNowTime()
		if err != nil {
			usecase.Logging.Error(err)
			return nil, err
		}

		outputs[i] = ports.FirestoreConnectOutputPort{
			Name:       doc.Data()["name"].(string),
			Address:    doc.Data()["address"].(string),
			Created_at: japaneseTime,
		}
	}
	return &outputs, nil
}

func (usecase *FirestoreConnectUsecase) Create(input *ports.FirestoreConnectInputPort) (*ports.FirestoreConnectOutputPort, error) {
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

	if firestoreConnect.Collection == "" {
		return nil, fmt.Errorf("BadRequest コレクションが入っていません。")
	}

	if firestoreConnect.Address == "" {
		return nil, fmt.Errorf("BadRequest 住所が入っていません。")
	}

	if firestoreConnect.Name == "" {
		return nil, fmt.Errorf("BadRequest 名前が入っていません。")
	}

	createdAt, err := usecase.FirestoreRepository.Insert(client, firestoreConnect)
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
