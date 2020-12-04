package ports

type TemperatureInputPort struct {
	Temp string `json:"temp" validate:"required" example:"23.1"` // 温度
	Humi string `json:"humi" validate:"required" example:"67.0"` // 湿度
}

type TemperatureOutputPort struct {
	ID   uint   `json:"id" example:"1"`
	Temp string `json:"temp" example:"27.5"` // 温度
	Humi string `json:"humi" example:"55.0"` // 湿度
}
