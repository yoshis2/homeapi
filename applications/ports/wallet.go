package ports

type WalletInputPort struct {
	Wallet string `json:"wallet" validate:"required" example:""`
}

type WalletOutputPort struct {
	ID     int
	Wallet string
}
