package ports

type TwitterInputPort struct {
	Tweet string `json:"tweet" validate:"required" example:"スリーネクストの投稿"`
}

type TwiterOutputPort struct {
	Tweet string
}
