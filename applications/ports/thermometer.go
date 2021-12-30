package ports

type ThermometerInputPort struct {
	Temperature string `json:"temperature" validate:"required" example:"23.1"` // 温度
	Humidity    string `json:"humidity" validate:"required" example:"67.0"`    // 湿度
}

type ThermometerOutputPort struct {
	ID          uint   `json:"id" example:"1"`
	Temperature string `json:"temperature" example:"27.5"` // 温度
	Humidity    string `json:"humidity" example:"55.0"`    // 湿度
}
