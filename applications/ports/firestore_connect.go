package ports

import "time"

type FirestoreConnectInputPort struct {
	Collection string `json:"collection" example:"members"`    // firestoreのコレクション
	Address    string `json:"address" example:"東京都港区六本木１−１−１"` // 名前
	Name       string `json:"name" example:"テスト　太郎"`           //  住所
}

type FirestoreConnectOutputPort struct {
	Address    string    `json:"address"` //  住所
	Name       string    `json:"name"`    // 名前
	Created_at time.Time `json:"created_at"`
}
