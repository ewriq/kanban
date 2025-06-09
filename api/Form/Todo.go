package Form

type  Todo struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Color    string `json:"color"`
	Token    string `json:"token"`
	User 	 string `json:"user"`
	Path string `json:"path"`
}

type TodoBody struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Color    string `json:"color"`
	User    string `json:"user"`
	Path string `json:"path"`
	Token    string `json:"token"`
}