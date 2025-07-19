package response

type CategoryResponse struct {
	Name string         `json:"name"`
	Buku []BukuResponse `json:"buku"`
}
