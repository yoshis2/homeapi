package ports

import "time"

type ImageInputPort struct {
	Name string `json:"name" validate:"required" example:"small phone"`
	Path string `json:"path" validate:"required" example:"/var/www/html"`
	Data string `json:"data" validate:"required" example:"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAhwAAAM2CAMAAACKYKqqAAAABGdBTUEAALGPC..."`
}

type ImageOutputPort struct {
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	CreatedAt time.Time `json:"created_at"`
}
