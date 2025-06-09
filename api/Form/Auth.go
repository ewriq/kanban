package Form

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	Perm 	 string `json:"perm"`
}

type UserBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserInfo struct {
	Token string `json:"token"`
}
