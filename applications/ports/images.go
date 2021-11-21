package ports

import "time"

type ImagesInputPort struct {
	ImageName string `json:"image_name" validate:"required" example:"small phone"`
	ImagePath string `json:"image_path" validate:"required" example:"/var/www/html"`
	ImageData string `json:"image_data" validate:"required" example:"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAhwAAAM2CAMAAACKYKqqAAAABGdBTUEAALGPC..."`
}

type ImagesOutputPort struct {
	ImageName string    `json:"image_name"`
	ImagePath string    `json:"image_path"`
	CreatedAt time.Time `json:"created_at"`
}
