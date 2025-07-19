package response

type BukuResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
	ReleaseYear int    `json:"releaseYear"`
	Price       int    `json:"price"`
	TotalPage   int    `json:"totalPage"`
	Thickness   string `json:"thickness"`
}
