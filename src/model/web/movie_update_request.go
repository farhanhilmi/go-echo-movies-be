package web

type MovieUpdateRequest struct {
	Id          int    `validate:"required" json:"id"`
	Title       string `validate:"required,min=1,max=255" json:"title"`
	Description string `validate:"required,min=1" json:"description"`
	Type        string `validate:"required,min=1,max=10" json:"type"`
	Imdb_url    string `validate:"required,min=1" json:"imdb_url"`
	Image_url   string `validate:"required,min=1" json:"image_url"`
	Watched     bool   `validate:"required,min=1" json:"watched"`
	Season      int    `json:"season"`
	Episode     int    `json:"episode"`
}
