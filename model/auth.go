package model

type Login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password"  validate:"required"`
}

type LoginResult struct {
	AccessToken string `json:"access_token"`
}
