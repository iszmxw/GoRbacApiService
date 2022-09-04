package requests

type Login struct {
	Username string `json:"username" validate:"username"`
	Password string `json:"password" validate:"password"`
}
