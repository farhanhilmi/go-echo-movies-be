package domain

type Movie struct {
	Id          int
	Title       string
	Description string
	Type        string
	Imdb_url    string
	Image_url   string
	Watched     bool
	Season      int
	Episode     int
}
