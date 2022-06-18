package models

type Token struct {
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
	Expire       int64       `json:"expire"`
	User         User        `json:"user"`
	Permissions  interface{} `json:"permissions"`
}
