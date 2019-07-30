package main

type News struct {
	Id        int  ` form:"id" json:"id"`
	Author   string `form:"author" json:"author"`
	Channel  	string `form:"channel" json:"channel"`
	Content  	string `form:"content" json:"content"`
	Person_id  	int `form:"person_id" json:"person_id"`

}
type Content struct{
	Content string
}
type NewsDTO struct {
	Id        int
	Author   string
	Channel  	string
	Content []string
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []News
}
