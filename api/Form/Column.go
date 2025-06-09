package Form
type  Column struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Color    string `json:"color"`
	Token    string `json:"token"`
	User 	 string `json:"user"`
}

type ColumnBody struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Color    string `json:"color"`
	User    string `json:"user"`
	Token    string `json:"id"`
}