package helpers

type Login struct {
	Name     string `json:"name" validator:"regexp=^[a-zA-Z0-9]$"`
	UserName string `json:"username" validator:"min=3,max=40,regexp=^[a-zA-Z0-9]$"`
	Passowrd string `json:"password" validator:"min=8"`
}
