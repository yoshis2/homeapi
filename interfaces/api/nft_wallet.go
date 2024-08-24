package api

import (
	"homeapi/applications/logging"

	"github.com/go-playground/validator/v10"
)

// FirestoreConnectController はFirestoreコネクト用コントローラー
type NftWalletController struct {
	// Usecase *usecases.FirestoreConnectUsecase
}

// NewFirestoreController はfirestoreコネクト用Newコントローラー
func NewNftWalletController(logging logging.Logging, validate *validator.Validate) *NftWalletController {
	return &NftWalletController{}
}
