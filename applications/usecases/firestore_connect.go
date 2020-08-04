package usecases

import (
	"net/http"

	"homeapi/domain"

	"homeapi/infrastructure/firebases"

	"homeapi/applications"
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

func (usecase *FirestoreConnectUsecase) List() (*[]ports.FirestoreConnectOutputPort, *applications.UsecaseError) {
	client, err := usecase.Firestore.Open()
	if usecaseError := applications.GetUErrorByError(err); usecaseError != nil {
		usecase.Logging.Error(usecaseError)
		return nil, usecaseError
	}

	iter := usecase.FirestoreRepository.List(client)
	docs, err := iter.GetAll()
	if usecaseError := applications.GetUErrorByError(err); usecaseError != nil {
		usecase.Logging.Error(usecaseError)
		return nil, usecaseError
	}

	outputs := make([]ports.FirestoreConnectOutputPort, len(docs))
	for i, doc := range docs {
		if doc.Data()["name"] == nil {
			usecaseError := &applications.UsecaseError{
				Code: http.StatusBadRequest,
				Msg:  "名前が存在しないデータがあります。",
			}
			return nil, usecaseError
		}

		if doc.Data()["address"] == nil {
			usecaseError := &applications.UsecaseError{
				Code: http.StatusBadRequest,
				Msg:  "住所が存在しないデータがあります。",
			}
			return nil, usecaseError
		}

		japaneseTime, err := util.JapaneseNowTime()
		if usecaseError := applications.GetUErrorByError(err); usecaseError != nil {
			usecase.Logging.Error(usecaseError)
			return nil, usecaseError
		}

		outputs[i] = ports.FirestoreConnectOutputPort{
			Name:       doc.Data()["name"].(string),
			Address:    doc.Data()["address"].(string),
			Created_at: japaneseTime,
		}
	}
	return &outputs, nil
}

func (usecase *FirestoreConnectUsecase) Create(input *ports.FirestoreConnectInputPort) (*ports.FirestoreConnectOutputPort, *applications.UsecaseError) {
	firestoreConnect := &domain.FirestoreConnect{
		Collection: input.Collection,
		Name:       input.Name,
		Address:    input.Address,
	}

	client, err := usecase.Firestore.Open()
	if usecaseError := applications.GetUErrorByError(err); usecaseError != nil {
		usecase.Logging.Error(usecaseError)
		return nil, usecaseError
	}

	if firestoreConnect.Collection == "" {
		usecaseError := &applications.UsecaseError{
			Code: http.StatusBadRequest,
			Msg:  "コレクションが入っていません。",
		}
		return nil, usecaseError
	}

	if firestoreConnect.Address == "" {
		usecaseError := &applications.UsecaseError{
			Code: http.StatusBadRequest,
			Msg:  "住所が入っていません。",
		}
		return nil, usecaseError
	}

	if firestoreConnect.Name == "" {
		usecaseError := &applications.UsecaseError{
			Code: http.StatusBadRequest,
			Msg:  "名前が入っていません。",
		}
		return nil, usecaseError
	}

	createdAt, err := usecase.FirestoreRepository.Insert(client, firestoreConnect)
	if usecaseError := applications.GetUErrorByError(err); usecaseError != nil {
		usecase.Logging.Error(usecaseError)
		return nil, usecaseError
	}

	defer client.Close()

	output := &ports.FirestoreConnectOutputPort{
		Name:       firestoreConnect.Name,
		Address:    firestoreConnect.Address,
		Created_at: createdAt.UpdateTime,
	}

	return output, nil
}
