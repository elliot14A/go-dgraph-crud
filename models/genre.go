package models

type Genere struct {
	Uid  string `json:"uid"`
	Name string `json:"name" validate:"min=1,max=50"`
}
