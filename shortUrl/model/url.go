package model

// URL 구조체
type Url struct {
	Id       int    `json:"id"`
	AliasURL string `json:"aliasUrl"`
	FullURL  string `json:"fullUrl"`
}
