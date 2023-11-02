package model

type Url struct {
	Id       int    `json:"id"`
	AliasURL string `json:"aliasUrl"`
	FullURL  string `json:"fullUrl"`
}
