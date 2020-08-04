package ports

import "time"

type ImagesInputPort struct {
	ImageName string `json:"image_name" example:"small phone"`
	ImagePath string `json:"image_path" example:"/var/www/html"`
	ImageData string `json:"image_data" example:"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAhwAAAM2CAMAAACKYKqqAAAABGdBTUEAALGPC..."`
}

type ImagesOutputPort struct {
	ImageName string
	ImagePath string
	CreatedAt time.Time
}
