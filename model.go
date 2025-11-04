package mockapi

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Category string `json:"category"`
	Desc     string `json:"desc"`
	CoverURL string `json:"cover_url"`
}
