package ports

type TwitterInputPort struct {
	Message string `json:"message" example:"スリーネクストの投稿"`
}

type TwiterOutputPort struct {
	Message string
}
