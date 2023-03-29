package web

type MovieResponse struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Imdb_url    string `json:"imdb_url"`
	Image_url   string `json:"image_url"`
	Watched     bool   `json:"watched"`
	Season      int    `json:"season"`
	Episode     int    `json:"episode"`
}
